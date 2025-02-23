const versionToken = "[VERSION]";

const unzip = require("unzipper");
const { Readable } = require("stream");
const { finished } = require("stream/promises");
const fs = require("fs");
const path = require("path");
const packageJson = require("../package.json");
const downloadedZipPath = path.join(
  __dirname,
  "..",
  "src",
  "openapi",
  "openapi-spec.zip",
);
const extractedZipPath = path.join(__dirname, "..", "src", "openapi");
const url = packageJson.custom.openapi.downloadUrl.replace(
  versionToken,
  packageJson.custom.openapi.version,
);

(async () => {
  try {
    if (fs.existsSync(extractedZipPath)) {
      console.info(`Cleaning up ${extractedZipPath}`);
      fs.rmSync(extractedZipPath, { recursive: true, force: true });
    }
    fs.mkdirSync(extractedZipPath);

    const response = await fetch(url);
    if (!response.ok) {
      console.error(`Failed to download openapi-spec from url: ${url}`);
      process.exit(1);
    }

    const file = fs.createWriteStream(downloadedZipPath);

    await finished(Readable.fromWeb(response.body).pipe(file));
    console.info(`Downloaded openapi-spec from ${url}`);

    await fs
      .createReadStream(downloadedZipPath)
      .pipe(unzip.Extract({ path: extractedZipPath }))
      .promise();
    console.info(`Extracting ${downloadedZipPath} to ${extractedZipPath}`);
  } catch (err) {
    console.error(`An error occurred when downloading ${url}`, err);
    process.exit(1);
  }
})();
