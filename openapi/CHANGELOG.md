# Changelog

## [1.10.1](https://github.com/rikotsev/markdown-blog/compare/api-v1.10.0...api-v1.10.1) (2025-05-09)


### Bug Fixes

* **openapi:** return the position field for page list ([#92](https://github.com/rikotsev/markdown-blog/issues/92)) ([37c08de](https://github.com/rikotsev/markdown-blog/commit/37c08de0219a3e6bc520c3d24232a993d823012f))

## [1.10.0](https://github.com/rikotsev/markdown-blog/compare/api-v1.9.0...api-v1.10.0) (2025-04-29)


### Features

* **openapi:** page position field ([#90](https://github.com/rikotsev/markdown-blog/issues/90)) ([2900607](https://github.com/rikotsev/markdown-blog/commit/29006078e4303429bcc2a9f8873232f67ca3e200))

## [1.9.0](https://github.com/rikotsev/markdown-blog/compare/api-v1.8.1...api-v1.9.0) (2025-03-09)


### Features

* **openapi:** CRUD for pages ([#86](https://github.com/rikotsev/markdown-blog/issues/86)) ([401c446](https://github.com/rikotsev/markdown-blog/commit/401c44617352e9ba800730aeb41826e190f088dd))

## [1.8.1](https://github.com/rikotsev/markdown-blog/compare/api-v1.8.0...api-v1.8.1) (2025-01-31)


### Bug Fixes

* **api:** included item as a separate model ([#64](https://github.com/rikotsev/markdown-blog/issues/64)) ([4baee07](https://github.com/rikotsev/markdown-blog/commit/4baee0736fd61dbc2957d2043388d9ed5e97515a))

## [1.8.0](https://github.com/rikotsev/markdown-blog/compare/api-v1.7.0...api-v1.8.0) (2025-01-25)


### Features

* **api, be, fe:** initial BE, improve API, generate stubs in FE ([#30](https://github.com/rikotsev/markdown-blog/issues/30)) ([102533d](https://github.com/rikotsev/markdown-blog/commit/102533d6d0cd9e5d593b401879726fd74d293f4f))
* **api:** compound documents implementation for article ([#50](https://github.com/rikotsev/markdown-blog/issues/50)) ([f286f52](https://github.com/rikotsev/markdown-blog/commit/f286f52c0a697043680bfd0d9b77287aff263462))
* **api:** discriminator field ([#52](https://github.com/rikotsev/markdown-blog/issues/52)) ([3603a1f](https://github.com/rikotsev/markdown-blog/commit/3603a1f977136e27a9991f731287a2b1243c1a03))
* **api:** json api includes and discriminator for entities ([#54](https://github.com/rikotsev/markdown-blog/issues/54)) ([8874227](https://github.com/rikotsev/markdown-blog/commit/887422778f6e1a0a3c31d99e61996af02ecf7a8b))
* **api:** publish to github pages ([#22](https://github.com/rikotsev/markdown-blog/issues/22)) ([88c0a4f](https://github.com/rikotsev/markdown-blog/commit/88c0a4fbb5c69a0f371163c24d6f895c47f26bc5))
* **api:** start using default problem+json ([#32](https://github.com/rikotsev/markdown-blog/issues/32)) ([82a75a1](https://github.com/rikotsev/markdown-blog/commit/82a75a1e54947ca056c6d74861662d209dc2c94d))
* **be, api:** setup category endpoints and add not found responses to api ([#42](https://github.com/rikotsev/markdown-blog/issues/42)) ([90f20ac](https://github.com/rikotsev/markdown-blog/commit/90f20ac15d85c0f0858cf8dc295135acfbc7c48c))
* CI/CD for markdown blog api ([#1](https://github.com/rikotsev/markdown-blog/issues/1)) ([ee8d6ae](https://github.com/rikotsev/markdown-blog/commit/ee8d6ae1ea0b06f8d714b2043a26d97a6d02147a))


### Bug Fixes

* **api:** added a proper description ([#6](https://github.com/rikotsev/markdown-blog/issues/6)) ([4c2b668](https://github.com/rikotsev/markdown-blog/commit/4c2b66821c648176680b22140cc485313b80e22d))
* **api:** authorized response for not protected paths ([#40](https://github.com/rikotsev/markdown-blog/issues/40)) ([f9c6d47](https://github.com/rikotsev/markdown-blog/commit/f9c6d47d5a8ffd4c9fb88b221608b2b2c0896a4e))
* **api:** content-type for category create conflict result ([#38](https://github.com/rikotsev/markdown-blog/issues/38)) ([d8c4fed](https://github.com/rikotsev/markdown-blog/commit/d8c4fedcba6285599dd02b1ce971a277cf09807e))
* **api:** markdown blog api release process ([#9](https://github.com/rikotsev/markdown-blog/issues/9)) ([c9c4c65](https://github.com/rikotsev/markdown-blog/commit/c9c4c6526b5383320e62e96a098ac20dc6ccad9a))
* **api:** proper setup of additional properties for problem+json ([#34](https://github.com/rikotsev/markdown-blog/issues/34)) ([129c971](https://github.com/rikotsev/markdown-blog/commit/129c9716875ad720b0dc08b07800ad66d6fe4aba))
* **api:** release flow 4 ([#18](https://github.com/rikotsev/markdown-blog/issues/18)) ([963ccdb](https://github.com/rikotsev/markdown-blog/commit/963ccdbfba5c313003f43de9846f9870c97aa586))
* **api:** release flow 5 ([#20](https://github.com/rikotsev/markdown-blog/issues/20)) ([7a17f0c](https://github.com/rikotsev/markdown-blog/commit/7a17f0c25dd9261c37ea15a7131a88cfd3df0547))
* **api:** revert to application/json because of generator issues ([#36](https://github.com/rikotsev/markdown-blog/issues/36)) ([b35b320](https://github.com/rikotsev/markdown-blog/commit/b35b320b2db0eeb29eae97adaf889b6354965a31))
* **api:** updated the docs ([#25](https://github.com/rikotsev/markdown-blog/issues/25)) ([9f66401](https://github.com/rikotsev/markdown-blog/commit/9f66401c31864849b82ca1e7c90970a3a52a2c1b))
* **api:** upload release spec ([#15](https://github.com/rikotsev/markdown-blog/issues/15)) ([1ee0fe6](https://github.com/rikotsev/markdown-blog/commit/1ee0fe6d459cde958d41b936284ae142f580bb60))
* **api:** versioning is reflected in the html ([#27](https://github.com/rikotsev/markdown-blog/issues/27)) ([05ba700](https://github.com/rikotsev/markdown-blog/commit/05ba7005bc1e6d9c96b11f40e3ab5ddb181b3ed6))
* openapi release again ([#11](https://github.com/rikotsev/markdown-blog/issues/11)) ([de63897](https://github.com/rikotsev/markdown-blog/commit/de638970c59a020211fc32edec9a022a39e53732))
* openapi release again 2 ([#13](https://github.com/rikotsev/markdown-blog/issues/13)) ([6d782c6](https://github.com/rikotsev/markdown-blog/commit/6d782c6960a2912bd182c79ea0d83c9fdc5354ad))

## [1.7.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.6.0...markdown-blog-api-v1.7.0) (2025-01-24)


### Features

* **api:** json api includes and discriminator for entities ([#54](https://github.com/rikotsev/markdown-blog/issues/54)) ([8874227](https://github.com/rikotsev/markdown-blog/commit/887422778f6e1a0a3c31d99e61996af02ecf7a8b))

## [1.6.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.5.0...markdown-blog-api-v1.6.0) (2024-12-23)


### Features

* **api:** discriminator field ([#52](https://github.com/rikotsev/markdown-blog/issues/52)) ([3603a1f](https://github.com/rikotsev/markdown-blog/commit/3603a1f977136e27a9991f731287a2b1243c1a03))

## [1.5.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.4.0...markdown-blog-api-v1.5.0) (2024-12-22)


### Features

* **api:** compound documents implementation for article ([#50](https://github.com/rikotsev/markdown-blog/issues/50)) ([f286f52](https://github.com/rikotsev/markdown-blog/commit/f286f52c0a697043680bfd0d9b77287aff263462))

## [1.4.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.3.4...markdown-blog-api-v1.4.0) (2024-11-12)


### Features

* **be, api:** setup category endpoints and add not found responses to api ([#42](https://github.com/rikotsev/markdown-blog/issues/42)) ([90f20ac](https://github.com/rikotsev/markdown-blog/commit/90f20ac15d85c0f0858cf8dc295135acfbc7c48c))

## [1.3.4](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.3.3...markdown-blog-api-v1.3.4) (2024-11-11)


### Bug Fixes

* **api:** authorized response for not protected paths ([#40](https://github.com/rikotsev/markdown-blog/issues/40)) ([f9c6d47](https://github.com/rikotsev/markdown-blog/commit/f9c6d47d5a8ffd4c9fb88b221608b2b2c0896a4e))

## [1.3.3](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.3.2...markdown-blog-api-v1.3.3) (2024-11-11)


### Bug Fixes

* **api:** content-type for category create conflict result ([#38](https://github.com/rikotsev/markdown-blog/issues/38)) ([d8c4fed](https://github.com/rikotsev/markdown-blog/commit/d8c4fedcba6285599dd02b1ce971a277cf09807e))

## [1.3.2](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.3.1...markdown-blog-api-v1.3.2) (2024-11-11)


### Bug Fixes

* **api:** revert to application/json because of generator issues ([#36](https://github.com/rikotsev/markdown-blog/issues/36)) ([b35b320](https://github.com/rikotsev/markdown-blog/commit/b35b320b2db0eeb29eae97adaf889b6354965a31))

## [1.3.1](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.3.0...markdown-blog-api-v1.3.1) (2024-11-11)


### Bug Fixes

* **api:** proper setup of additional properties for problem+json ([#34](https://github.com/rikotsev/markdown-blog/issues/34)) ([129c971](https://github.com/rikotsev/markdown-blog/commit/129c9716875ad720b0dc08b07800ad66d6fe4aba))

## [1.3.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.2.0...markdown-blog-api-v1.3.0) (2024-11-11)


### Features

* **api:** start using default problem+json ([#32](https://github.com/rikotsev/markdown-blog/issues/32)) ([82a75a1](https://github.com/rikotsev/markdown-blog/commit/82a75a1e54947ca056c6d74861662d209dc2c94d))

## [1.2.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.1.2...markdown-blog-api-v1.2.0) (2024-11-09)


### Features

* **api, be, fe:** initial BE, improve API, generate stubs in FE ([#30](https://github.com/rikotsev/markdown-blog/issues/30)) ([102533d](https://github.com/rikotsev/markdown-blog/commit/102533d6d0cd9e5d593b401879726fd74d293f4f))

## [1.1.2](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.1.1...markdown-blog-api-v1.1.2) (2024-10-12)


### Bug Fixes

* **api:** versioning is reflected in the html ([#27](https://github.com/rikotsev/markdown-blog/issues/27)) ([05ba700](https://github.com/rikotsev/markdown-blog/commit/05ba7005bc1e6d9c96b11f40e3ab5ddb181b3ed6))

## [1.1.1](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.1.0...markdown-blog-api-v1.1.1) (2024-10-12)


### Bug Fixes

* **api:** updated the docs ([#25](https://github.com/rikotsev/markdown-blog/issues/25)) ([9f66401](https://github.com/rikotsev/markdown-blog/commit/9f66401c31864849b82ca1e7c90970a3a52a2c1b))

## [1.1.0](https://github.com/rikotsev/markdown-blog/compare/markdown-blog-api-v1.0.0...markdown-blog-api-v1.1.0) (2024-10-12)


### Features

* **api:** publish to github pages ([#22](https://github.com/rikotsev/markdown-blog/issues/22)) ([88c0a4f](https://github.com/rikotsev/markdown-blog/commit/88c0a4fbb5c69a0f371163c24d6f895c47f26bc5))

## 1.0.0 (2024-10-10)


### Features

* CI/CD for markdown blog api ([#1](https://github.com/rikotsev/markdown-blog/issues/1)) ([ee8d6ae](https://github.com/rikotsev/markdown-blog/commit/ee8d6ae1ea0b06f8d714b2043a26d97a6d02147a))


### Bug Fixes

* **api:** added a proper description ([#6](https://github.com/rikotsev/markdown-blog/issues/6)) ([4c2b668](https://github.com/rikotsev/markdown-blog/commit/4c2b66821c648176680b22140cc485313b80e22d))
* **api:** markdown blog api release process ([#9](https://github.com/rikotsev/markdown-blog/issues/9)) ([c9c4c65](https://github.com/rikotsev/markdown-blog/commit/c9c4c6526b5383320e62e96a098ac20dc6ccad9a))
* **api:** release flow 4 ([#18](https://github.com/rikotsev/markdown-blog/issues/18)) ([963ccdb](https://github.com/rikotsev/markdown-blog/commit/963ccdbfba5c313003f43de9846f9870c97aa586))
* **api:** release flow 5 ([#20](https://github.com/rikotsev/markdown-blog/issues/20)) ([7a17f0c](https://github.com/rikotsev/markdown-blog/commit/7a17f0c25dd9261c37ea15a7131a88cfd3df0547))
* **api:** upload release spec ([#15](https://github.com/rikotsev/markdown-blog/issues/15)) ([1ee0fe6](https://github.com/rikotsev/markdown-blog/commit/1ee0fe6d459cde958d41b936284ae142f580bb60))
* openapi release again ([#11](https://github.com/rikotsev/markdown-blog/issues/11)) ([de63897](https://github.com/rikotsev/markdown-blog/commit/de638970c59a020211fc32edec9a022a39e53732))
* openapi release again 2 ([#13](https://github.com/rikotsev/markdown-blog/issues/13)) ([6d782c6](https://github.com/rikotsev/markdown-blog/commit/6d782c6960a2912bd182c79ea0d83c9fdc5354ad))
