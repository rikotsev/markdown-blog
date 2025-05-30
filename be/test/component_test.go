package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rikotsev/markdown-blog/be/internal/config"
	"github.com/rikotsev/markdown-blog/be/internal/server"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

type ApplicationSuite struct {
	suite.Suite
	ctx                context.Context
	cfg                *config.Config
	dbPoolForTests     *pgxpool.Pool
	applicationContext server.ApplicationContext
	applicationServer  server.ApplicationServer
	serverAddr         string
	dockerComposeDown  func()
	httpClient         *http.Client
	testSchemaCounter  int
	testSchema         string
}

type TestContext struct {
	ctx  context.Context
	cfg  *config.Config
	pool *pgxpool.Pool
}

func (t *TestContext) Ctx() context.Context {
	return t.ctx
}

func (t *TestContext) Cfg() *config.Config {
	return t.cfg
}

func (t *TestContext) Pool() *pgxpool.Pool {
	return t.pool
}

const validBearerToken = "Bearer mock_token"

type TestAuthProvider struct {
	tokens map[string]bool
}

func (p *TestAuthProvider) Handle(r *http.Request) (server.Identity, error) {
	token := r.Header.Get("Authorization")
	if valid, found := p.tokens[token]; !found || !valid {
		return server.Identity{}, server.ErrUnauthorized
	}

	return server.Identity{
		Email: "john.doe@example.com",
	}, nil
}

func (s *ApplicationSuite) SetupSuite() {
	s.ctx = context.Background()
	cfg, err := config.InitConfig()
	s.NoError(err)
	s.cfg = cfg

	s.startDb()
	s.startServer()
	s.httpClient = &http.Client{}
}

func (s *ApplicationSuite) TearDownSuite() {
	if s.dockerComposeDown != nil {
		s.dockerComposeDown()
	}
}

func (s *ApplicationSuite) startServer() {
	testContext := TestContext{
		ctx: s.ctx,
		cfg: s.cfg,
	}

	dbCfg, err := pgxpool.ParseConfig(s.cfg.Database.Url)
	s.Require().NoError(err, "could not create db configuration for url", s.cfg.Database.Url)
	dbCfg.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
		_, err := conn.Exec(ctx, fmt.Sprintf("SET search_path TO %s", s.testSchema))
		return err == nil
	}
	pool, err := pgxpool.NewWithConfig(s.ctx, dbCfg)
	s.Require().NoError(err, "could not create db pool")
	testContext.pool = pool

	testAuthProvider := TestAuthProvider{
		tokens: map[string]bool{
			validBearerToken: true,
		},
	}

	s.applicationContext = &testContext
	srv, err := server.New(&testContext, &testAuthProvider)
	s.Require().NoError(err)
	go func() {
		err := srv.Start()
		s.NoError(err, "failed to start listening")
	}()
	s.applicationServer = srv
	s.serverAddr = fmt.Sprintf("localhost:%d", srv.Listener.Addr().(*net.TCPAddr).Port)
}

func (s *ApplicationSuite) startDb() {
	compose, err := tc.NewDockerCompose("../docker/docker-compose-test.yaml")
	s.Require().NoError(err, "could not find docker-compose file")
	s.dockerComposeDown = func() {
		ctx, cancelFunc := context.WithTimeout(s.ctx, time.Second*10)
		defer cancelFunc()
		err := compose.Down(ctx, tc.RemoveOrphans(true), tc.RemoveImagesLocal)
		s.Require().NoError(err, "could not execute successfully docker-compose down")
	}

	ctx, cancelFunc := context.WithTimeout(s.ctx, time.Minute)
	defer cancelFunc()
	err = compose.Up(ctx, tc.Wait(true))
	if err != nil && strings.Contains(err.Error(), "golang-migrate-1 exited (0)") {
		pool, err := pgxpool.New(s.ctx, s.cfg.Database.Url)
		s.Require().NoError(err, "could not set-up a db connection pool for setting up tests")
		s.dbPoolForTests = pool
		return
	}
	s.Require().NoError(err, "could not execute successfully docker-compose up")
}

func (s *ApplicationSuite) SetupTest() {
	s.testSchemaCounter++
	s.testSchema = fmt.Sprintf("test_schema_%d", s.testSchemaCounter)
	_, err := s.dbPoolForTests.Exec(s.ctx, fmt.Sprintf("CREATE SCHEMA %s;", s.testSchema))
	s.Require().NoError(err)
	m, err := migrate.New("file://../db", fmt.Sprintf("%s&search_path=%s", s.cfg.Database.Url, s.testSchema))
	s.Require().NoError(err)
	err = m.Up()
	s.Require().NoError(err)
}

func (s *ApplicationSuite) httpGet(path string, result any) {
	url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	s.Require().NoError(err, "failed to create request", url)

	resp, err := s.httpClient.Do(req)
	s.Require().NoError(err, "failed to execute request", url, req)

	content, err := io.ReadAll(resp.Body)
	s.Require().NoError(err, "could not resp body", url, req, resp)

	err = json.Unmarshal(content, result)
	s.Require().NoError(err, "could not unmarshal resp body", url, req, string(content))
}

func (s *ApplicationSuite) httpDelete(path string) int {
	//url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)
	//
	//req, err := http.NewRequest(http.MethodDelete, url, nil)
	//s.Require().NoError(err, "failed to create request", url)
	//req.Header.Set("Authorization", validBearerToken)
	//req.Header.Set("Content-Type", "application/json")
	//
	//resp, err := s.httpClient.Do(req)
	//s.Require().NoError(err, "failed to execute request", url, req)
	//
	//return resp.StatusCode

	resp := s.httpDeleteWithHeaders(path, map[string]string{
		"Authorization": validBearerToken,
		"Content-Type":  "application/json",
	})

	return resp.StatusCode
}

func (s *ApplicationSuite) httpDeleteWithHeaders(path string, headers map[string]string) *http.Response {
	url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	s.Require().NoError(err, "failed to create request", url)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := s.httpClient.Do(req)
	s.Require().NoError(err, "failed to execute request", url, req)

	return resp
}

func (s *ApplicationSuite) httpPost(path string, payload any, result any) int {
	url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)

	requestContent, err := json.Marshal(payload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestContent))
	s.Require().NoError(err, "failed to create request ", url)
	req.Header.Set("Authorization", validBearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	s.Require().NoError(err, "failed to execute request", url, req)

	content, err := io.ReadAll(resp.Body)
	s.Require().NoError(err, "could not read resp body", url)

	err = json.Unmarshal(content, result)
	s.Require().NoError(err, "could not unmarshal resp body", url, string(content))

	return resp.StatusCode
}

func (s *ApplicationSuite) httpPostRaw(path string, payload any) *http.Response {
	return s.httpPostRawWithHeaders(path, payload, map[string]string{
		"Authorization": validBearerToken,
		"Content-Type":  "application/json",
	})
}

func (s *ApplicationSuite) httpPostRawWithHeaders(path string, payload any, headers map[string]string) *http.Response {
	url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)

	requestContent, err := json.Marshal(payload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestContent))
	s.Require().NoError(err, "failed to create request ", url)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := s.httpClient.Do(req)
	s.Require().NoError(err, "failed to execute request", url, req)

	return resp
}

func (s *ApplicationSuite) httpPatchRaw(path string, payload any) *http.Response {
	return s.httpPatchRawWithHeaders(path, payload, map[string]string{
		"Authorization": validBearerToken,
		"Content-Type":  "application/json",
	})
}

func (s *ApplicationSuite) httpPatchRawWithHeaders(path string, payload any, headers map[string]string) *http.Response {
	url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)

	requestContent, err := json.Marshal(payload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(requestContent))
	s.Require().NoError(err, "failed to create request ", url)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := s.httpClient.Do(req)
	s.Require().NoError(err, "failed to execute request", url, req)

	return resp
}

func (s *ApplicationSuite) httpGetRaw(path string) *http.Response {
	url := fmt.Sprintf("http://%s/%s", s.serverAddr, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	s.Require().NoError(err, "failed to create request", url)

	resp, err := s.httpClient.Do(req)
	s.Require().NoError(err, "failed to execute request", url, req)

	return resp
}

func TestMarkdownBlogSuite(t *testing.T) {
	suite.Run(t, new(ApplicationSuite))
}

func ptr[T any](v T) *T { return &v }
