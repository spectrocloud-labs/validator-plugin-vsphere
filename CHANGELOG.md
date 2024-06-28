# Changelog

## [0.0.27](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.26...v0.0.27) (2024-06-28)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 7f521ea ([#285](https://github.com/validator-labs/validator-plugin-vsphere/issues/285)) ([598a86f](https://github.com/validator-labs/validator-plugin-vsphere/commit/598a86fc33e4784225661204a7df7ba3d39ddc9a))
* **deps:** update module github.com/validator-labs/validator to v0.0.43 ([#287](https://github.com/validator-labs/validator-plugin-vsphere/issues/287)) ([49e45e9](https://github.com/validator-labs/validator-plugin-vsphere/commit/49e45e9a790b27d115ca11721897c1e7bb716d35))
* **deps:** update module github.com/vmware/govmomi to v0.38.0 ([#288](https://github.com/validator-labs/validator-plugin-vsphere/issues/288)) ([8a98b24](https://github.com/validator-labs/validator-plugin-vsphere/commit/8a98b24c208576517a4c0cde129a1c3bc5a34686))

## [0.0.26](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.25...v0.0.26) (2024-06-12)


### Other

* ensure that VSphereCloudDriver has IsAdminAccount function ([#284](https://github.com/validator-labs/validator-plugin-vsphere/issues/284)) ([ad92e4a](https://github.com/validator-labs/validator-plugin-vsphere/commit/ad92e4a9ced5c5f4aa175afc5bd7bf9c67657f5a))
* ensure that VSphereCloudDriver implements the VsphereDriver interface ([#282](https://github.com/validator-labs/validator-plugin-vsphere/issues/282)) ([6ff5e3b](https://github.com/validator-labs/validator-plugin-vsphere/commit/6ff5e3b9281b2cb30cd2c9c57e2a841c0086d54f))

## [0.0.25](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.24...v0.0.25) (2024-06-12)


### Bug Fixes

* **deps:** update kubernetes packages to v0.30.2 ([#280](https://github.com/validator-labs/validator-plugin-vsphere/issues/280)) ([d262e44](https://github.com/validator-labs/validator-plugin-vsphere/commit/d262e44e1fa39140b91441b322890ab4ab8e3538))
* **deps:** update module sigs.k8s.io/cluster-api to v1.7.3 ([#278](https://github.com/validator-labs/validator-plugin-vsphere/issues/278)) ([c5e13f8](https://github.com/validator-labs/validator-plugin-vsphere/commit/c5e13f8ed5d37130ef96c63e6b31242ee3e34fe1))


### Refactoring

* move IsAdmin out from internal package ([#281](https://github.com/validator-labs/validator-plugin-vsphere/issues/281)) ([84ed6c4](https://github.com/validator-labs/validator-plugin-vsphere/commit/84ed6c4c762ae53ceac1d94c8e73dde86445ba79))

## [0.0.24](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.23...v0.0.24) (2024-06-09)


### Bug Fixes

* add yaml struct tags to VsphereCloudAccount ([#276](https://github.com/validator-labs/validator-plugin-vsphere/issues/276)) ([923a14c](https://github.com/validator-labs/validator-plugin-vsphere/commit/923a14c76dc22938f3006840625a51deee7b7c27))

## [0.0.23](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.22...v0.0.23) (2024-06-07)


### Features

* expose new functions in vsphere driver ([#272](https://github.com/validator-labs/validator-plugin-vsphere/issues/272)) ([f3dbf24](https://github.com/validator-labs/validator-plugin-vsphere/commit/f3dbf24aae310846792fa83072b5c23cc3e99fca))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 404ba88 ([#260](https://github.com/validator-labs/validator-plugin-vsphere/issues/260)) ([a571144](https://github.com/validator-labs/validator-plugin-vsphere/commit/a5711447dfb7e431c85f4bb882e652891b943dce))
* **deps:** update golang.org/x/exp digest to fc45aab ([#267](https://github.com/validator-labs/validator-plugin-vsphere/issues/267)) ([7249aca](https://github.com/validator-labs/validator-plugin-vsphere/commit/7249aca45071a710942d7f04ad49fcfb5cb791ea))
* **deps:** update golang.org/x/exp digest to fd00a4e ([#266](https://github.com/validator-labs/validator-plugin-vsphere/issues/266)) ([ae890f7](https://github.com/validator-labs/validator-plugin-vsphere/commit/ae890f7aa6f111c5089a68cc454a9daba36f8dcc))
* **deps:** update module github.com/go-logr/logr to v1.4.2 ([#261](https://github.com/validator-labs/validator-plugin-vsphere/issues/261)) ([9ccc44a](https://github.com/validator-labs/validator-plugin-vsphere/commit/9ccc44a20a3d01522a8fa9a7ec27bc9750b783ad))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.17.3 ([#237](https://github.com/validator-labs/validator-plugin-vsphere/issues/237)) ([44917d6](https://github.com/validator-labs/validator-plugin-vsphere/commit/44917d600369b596af39bdb1462a204bece0fed3))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.19.0 ([#256](https://github.com/validator-labs/validator-plugin-vsphere/issues/256)) ([50502df](https://github.com/validator-labs/validator-plugin-vsphere/commit/50502df6c98277c68c4f94ad142eb562b34ef87c))
* **deps:** update module github.com/validator-labs/validator to v0.0.41 ([#263](https://github.com/validator-labs/validator-plugin-vsphere/issues/263)) ([d977189](https://github.com/validator-labs/validator-plugin-vsphere/commit/d9771892ffb53e85d8db2840655fc5fb72103128))
* **deps:** update module github.com/validator-labs/validator to v0.0.42 ([#274](https://github.com/validator-labs/validator-plugin-vsphere/issues/274)) ([1a06968](https://github.com/validator-labs/validator-plugin-vsphere/commit/1a06968a121c8672a8f8f45b0f97fdc9fdb64ddf))
* **deps:** update module github.com/vmware/govmomi to v0.37.2 ([#238](https://github.com/validator-labs/validator-plugin-vsphere/issues/238)) ([32b5dc6](https://github.com/validator-labs/validator-plugin-vsphere/commit/32b5dc6e4e72469862b55f9dedf5640c9e6af21d))
* **deps:** update module github.com/vmware/govmomi to v0.37.3 ([#269](https://github.com/validator-labs/validator-plugin-vsphere/issues/269)) ([3aca407](https://github.com/validator-labs/validator-plugin-vsphere/commit/3aca4072400730c8737a14fe034a1b8483f6174a))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.18.2 ([#241](https://github.com/validator-labs/validator-plugin-vsphere/issues/241)) ([98b9f19](https://github.com/validator-labs/validator-plugin-vsphere/commit/98b9f194b43869c30e5248807b84a361c0f1b528))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.18.4 ([#271](https://github.com/validator-labs/validator-plugin-vsphere/issues/271)) ([ff2794b](https://github.com/validator-labs/validator-plugin-vsphere/commit/ff2794be13125b6b4165d653302aba6c78f41ef5))


### Other

* **deps:** bump golang.org/x/net from 0.22.0 to 0.23.0 ([#242](https://github.com/validator-labs/validator-plugin-vsphere/issues/242)) ([e0858b0](https://github.com/validator-labs/validator-plugin-vsphere/commit/e0858b0fbfb2b767b5af29bb723b20007dae8a08))
* **deps:** pin googleapis/release-please-action action to f3969c0 ([#252](https://github.com/validator-labs/validator-plugin-vsphere/issues/252)) ([ce3c62f](https://github.com/validator-labs/validator-plugin-vsphere/commit/ce3c62f026eb618e33ca0f3718ff4456ca90a348))
* **deps:** update actions/checkout digest to 0ad4b8f ([#244](https://github.com/validator-labs/validator-plugin-vsphere/issues/244)) ([340f1a2](https://github.com/validator-labs/validator-plugin-vsphere/commit/340f1a2812a9ecaa79b7f86429bd5b41562205d7))
* **deps:** update actions/checkout digest to a5ac7e5 ([#253](https://github.com/validator-labs/validator-plugin-vsphere/issues/253)) ([cc074ee](https://github.com/validator-labs/validator-plugin-vsphere/commit/cc074ee7eacbc02ceec1b63f700a786ebe40e712))
* **deps:** update actions/setup-go digest to cdcb360 ([#257](https://github.com/validator-labs/validator-plugin-vsphere/issues/257)) ([6b267d5](https://github.com/validator-labs/validator-plugin-vsphere/commit/6b267d52327661f16c5eebf26202a6c25b82c9d0))
* **deps:** update anchore/sbom-action action to v0.15.11 ([#236](https://github.com/validator-labs/validator-plugin-vsphere/issues/236)) ([87e92ca](https://github.com/validator-labs/validator-plugin-vsphere/commit/87e92cae74ded9412dc796bb9988a95c2602f36b))
* **deps:** update anchore/sbom-action action to v0.16.0 ([#255](https://github.com/validator-labs/validator-plugin-vsphere/issues/255)) ([ee50eb9](https://github.com/validator-labs/validator-plugin-vsphere/commit/ee50eb9882a27d39ec0e4c8f8efcecc8234eb98c))
* **deps:** update azure/setup-helm digest to fe7b79c ([#243](https://github.com/validator-labs/validator-plugin-vsphere/issues/243)) ([b01c2bd](https://github.com/validator-labs/validator-plugin-vsphere/commit/b01c2bda1ac63c6f2c0893fb75fba9479b9ebc23))
* **deps:** update codecov/codecov-action digest to 125fc84 ([#254](https://github.com/validator-labs/validator-plugin-vsphere/issues/254)) ([f2e7f45](https://github.com/validator-labs/validator-plugin-vsphere/commit/f2e7f454f79f8552dbef819ec1bfa5ee86831497))
* **deps:** update codecov/codecov-action digest to 6d79887 ([#234](https://github.com/validator-labs/validator-plugin-vsphere/issues/234)) ([264c4c2](https://github.com/validator-labs/validator-plugin-vsphere/commit/264c4c205656b044ba397ccbd929d949655fff90))
* **deps:** update dependency go to v1.22.4 ([#268](https://github.com/validator-labs/validator-plugin-vsphere/issues/268)) ([045458e](https://github.com/validator-labs/validator-plugin-vsphere/commit/045458ecc3f0a731f0e931656c7bcaedaa6e0679))
* **deps:** update docker/login-action digest to 0d4c9c5 ([#258](https://github.com/validator-labs/validator-plugin-vsphere/issues/258)) ([1fbdb5e](https://github.com/validator-labs/validator-plugin-vsphere/commit/1fbdb5e6706d6d6871dbc608742cea9c6f329203))
* **deps:** update docker/setup-buildx-action digest to d70bba7 ([#240](https://github.com/validator-labs/validator-plugin-vsphere/issues/240)) ([5557060](https://github.com/validator-labs/validator-plugin-vsphere/commit/555706061b043090568367853b1e0611bfaaca01))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.16.0 ([#239](https://github.com/validator-labs/validator-plugin-vsphere/issues/239)) ([205d85b](https://github.com/validator-labs/validator-plugin-vsphere/commit/205d85be5fb286eeb27b9cee5b91c77ff3269f19))
* **deps:** update helm/kind-action action to v1.10.0 ([#264](https://github.com/validator-labs/validator-plugin-vsphere/issues/264)) ([9663041](https://github.com/validator-labs/validator-plugin-vsphere/commit/96630418bf612d6c18fc5848ed03dd0868c7ae2e))
* **deps:** update softprops/action-gh-release digest to 69320db ([#259](https://github.com/validator-labs/validator-plugin-vsphere/issues/259)) ([80b1725](https://github.com/validator-labs/validator-plugin-vsphere/commit/80b1725284d48814a5dc1c8f50d508cb5ffbba95))
* release 0.0.23 ([a91d7b3](https://github.com/validator-labs/validator-plugin-vsphere/commit/a91d7b34e3cf26b1292d15014bb5931a730f422e))

## [0.0.22](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.21...v0.0.22) (2024-05-28)


### Features

* add support for validating role privileges on your own non admin account ([#249](https://github.com/validator-labs/validator-plugin-vsphere/issues/249)) ([ca4a80a](https://github.com/validator-labs/validator-plugin-vsphere/commit/ca4a80ac06930e94e9aacdb3ee9212d1f8589c34))

## [0.0.21](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.20...v0.0.21) (2024-05-23)


### Other

* ensure validation results are updated when role privilege rules fail ([#247](https://github.com/validator-labs/validator-plugin-vsphere/issues/247)) ([cd77e69](https://github.com/validator-labs/validator-plugin-vsphere/commit/cd77e69dc48e0695eb15b129a4bb157cef07dec4))
* setup devspace ([#246](https://github.com/validator-labs/validator-plugin-vsphere/issues/246)) ([e384de2](https://github.com/validator-labs/validator-plugin-vsphere/commit/e384de25d33ea2969ca429db61f1974b455c4def))

## [0.0.20](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.19...v0.0.20) (2024-05-17)


### Bug Fixes

* **deps:** update github.com/spectrocloud-labs/validator digest to fc351f3 ([#220](https://github.com/validator-labs/validator-plugin-vsphere/issues/220)) ([f490192](https://github.com/validator-labs/validator-plugin-vsphere/commit/f490192862c6a7ccd8d8ccaea9f90f8840bd5270))
* **deps:** update golang.org/x/exp digest to a85f2c6 ([#230](https://github.com/validator-labs/validator-plugin-vsphere/issues/230)) ([53ed804](https://github.com/validator-labs/validator-plugin-vsphere/commit/53ed804b123527981aaadac0a95e2b0329ebfd25))
* **deps:** update golang.org/x/exp digest to c7f7c64 ([#228](https://github.com/validator-labs/validator-plugin-vsphere/issues/228)) ([99b5d32](https://github.com/validator-labs/validator-plugin-vsphere/commit/99b5d321ff6dbadbe2d9185f8e7b9b713210a3e8))
* **deps:** update kubernetes packages to v0.29.3 ([#229](https://github.com/validator-labs/validator-plugin-vsphere/issues/229)) ([217bf6c](https://github.com/validator-labs/validator-plugin-vsphere/commit/217bf6c9a81dfd8c44bb333863b45bc27a11a5dd))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.17.0 ([#231](https://github.com/validator-labs/validator-plugin-vsphere/issues/231)) ([4175b26](https://github.com/validator-labs/validator-plugin-vsphere/commit/4175b2695b2b8b553accf3f9b265fe90c2642bb5))
* **deps:** update module github.com/onsi/gomega to v1.32.0 ([#232](https://github.com/validator-labs/validator-plugin-vsphere/issues/232)) ([0f0daac](https://github.com/validator-labs/validator-plugin-vsphere/commit/0f0daac1dd64647941261f65bbef49cc5215fcda))
* **deps:** update module github.com/spectrocloud-labs/validator to v0.0.38 ([#184](https://github.com/validator-labs/validator-plugin-vsphere/issues/184)) ([e692268](https://github.com/validator-labs/validator-plugin-vsphere/commit/e692268aa2fbe473ff4c114d8a08663df163a1f3))
* **deps:** update module github.com/vmware/govmomi to v0.36.1 ([#218](https://github.com/validator-labs/validator-plugin-vsphere/issues/218)) ([8f85fd6](https://github.com/validator-labs/validator-plugin-vsphere/commit/8f85fd6513c2a6d5adff8d947049a61f317eb371))
* **deps:** update module sigs.k8s.io/cluster-api to v1.6.3 ([#221](https://github.com/validator-labs/validator-plugin-vsphere/issues/221)) ([c575d8b](https://github.com/validator-labs/validator-plugin-vsphere/commit/c575d8bb9943c9bb6ff5847182b0deb300af6dbe))


### Other

* **deps:** update actions/setup-python digest to 82c7e63 ([#233](https://github.com/validator-labs/validator-plugin-vsphere/issues/233)) ([633db14](https://github.com/validator-labs/validator-plugin-vsphere/commit/633db1457862d130ba37c8263c3820b3faee523d))
* **deps:** update docker/build-push-action digest to 2cdde99 ([#225](https://github.com/validator-labs/validator-plugin-vsphere/issues/225)) ([be0eff0](https://github.com/validator-labs/validator-plugin-vsphere/commit/be0eff0dc652b1dfb398a2c89b352c3e5db6f329))
* **deps:** update docker/login-action digest to e92390c ([#222](https://github.com/validator-labs/validator-plugin-vsphere/issues/222)) ([5a02db5](https://github.com/validator-labs/validator-plugin-vsphere/commit/5a02db50b22a0bc72c8f40f4a10d49fa298db315))
* **deps:** update docker/setup-buildx-action digest to 2b51285 ([#226](https://github.com/validator-labs/validator-plugin-vsphere/issues/226)) ([c531b2d](https://github.com/validator-labs/validator-plugin-vsphere/commit/c531b2ddec2af071787e727d62e15aa5b5c19497))
* **deps:** update softprops/action-gh-release digest to 9d7c94c ([#217](https://github.com/validator-labs/validator-plugin-vsphere/issues/217)) ([5e1b282](https://github.com/validator-labs/validator-plugin-vsphere/commit/5e1b28292d26d23996e2fddb698555c8d63d1a29))
* migrate from spectrocloud-labs to validator-labs ([#245](https://github.com/validator-labs/validator-plugin-vsphere/issues/245)) ([34a42da](https://github.com/validator-labs/validator-plugin-vsphere/commit/34a42da08950a54bd427db578667216178d2ddae))

## [0.0.19](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.18...v0.0.19) (2024-03-13)


### Other

* **deps:** update google-github-actions/release-please-action digest to a37ac6e ([#214](https://github.com/validator-labs/validator-plugin-vsphere/issues/214)) ([d038990](https://github.com/validator-labs/validator-plugin-vsphere/commit/d0389903f47d3fd9ce4b0f82efeaeac63ab6aa5e))
* **deps:** update softprops/action-gh-release digest to 3198ee1 ([#213](https://github.com/validator-labs/validator-plugin-vsphere/issues/213)) ([77f1867](https://github.com/validator-labs/validator-plugin-vsphere/commit/77f18676653a02b99cc19510b7bdbe13fa87433d))


### Refactoring

* use patch helpers ([#219](https://github.com/validator-labs/validator-plugin-vsphere/issues/219)) ([5f3b8db](https://github.com/validator-labs/validator-plugin-vsphere/commit/5f3b8db4e8b3fbec6fc7d5f6b8ce9a991def7c5b))

## [0.0.18](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.17...v0.0.18) (2024-03-11)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 2c58cdc ([#186](https://github.com/validator-labs/validator-plugin-vsphere/issues/186)) ([8567356](https://github.com/validator-labs/validator-plugin-vsphere/commit/8567356b68f34fcbda309ceabde2a6611bfb983d))
* **deps:** update golang.org/x/exp digest to 814bf88 ([#199](https://github.com/validator-labs/validator-plugin-vsphere/issues/199)) ([beb122d](https://github.com/validator-labs/validator-plugin-vsphere/commit/beb122d5ddd0886fc70d4da6164c17eef8e0b3b0))
* **deps:** update golang.org/x/exp digest to ec58324 ([#193](https://github.com/validator-labs/validator-plugin-vsphere/issues/193)) ([48829e4](https://github.com/validator-labs/validator-plugin-vsphere/commit/48829e4a85f93d0a1579a7d26c3e64eca6a41747))
* **deps:** update kubernetes packages to v0.29.2 ([#194](https://github.com/validator-labs/validator-plugin-vsphere/issues/194)) ([75c6f8d](https://github.com/validator-labs/validator-plugin-vsphere/commit/75c6f8de43cf18be1708f052180727e158d705de))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.16.0 ([#205](https://github.com/validator-labs/validator-plugin-vsphere/issues/205)) ([994993d](https://github.com/validator-labs/validator-plugin-vsphere/commit/994993d05c9112d71b0f65178ccd3503ccf8df79))
* **deps:** update module github.com/vmware/govmomi to v0.35.0 ([#191](https://github.com/validator-labs/validator-plugin-vsphere/issues/191)) ([44873bc](https://github.com/validator-labs/validator-plugin-vsphere/commit/44873bc03063a78ca29c38df9ae40cdc105f8ef8))
* **deps:** update module github.com/vmware/govmomi to v0.36.0 ([#207](https://github.com/validator-labs/validator-plugin-vsphere/issues/207)) ([69b320a](https://github.com/validator-labs/validator-plugin-vsphere/commit/69b320a9751ad310f2bc4d5ffc99d14864454196))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.17.1 ([#163](https://github.com/validator-labs/validator-plugin-vsphere/issues/163)) ([cb12fa0](https://github.com/validator-labs/validator-plugin-vsphere/commit/cb12fa087486924314513c4fbca90d0f3d10516b))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.17.2 ([#198](https://github.com/validator-labs/validator-plugin-vsphere/issues/198)) ([eaf76fa](https://github.com/validator-labs/validator-plugin-vsphere/commit/eaf76fae61c31e5707f9f544df4e7180b06a718d))
* err handling in CheckTestCase ([#212](https://github.com/validator-labs/validator-plugin-vsphere/issues/212)) ([de298c3](https://github.com/validator-labs/validator-plugin-vsphere/commit/de298c3d4092e08ed65d52cfcbb27227228a5007))


### Other

* **deps:** update actions/upload-artifact digest to 5d5d22a ([#187](https://github.com/validator-labs/validator-plugin-vsphere/issues/187)) ([67a9d28](https://github.com/validator-labs/validator-plugin-vsphere/commit/67a9d28afd5fe3f1d78d92e24413c48f4d051df5))
* **deps:** update anchore/sbom-action action to v0.15.9 ([#206](https://github.com/validator-labs/validator-plugin-vsphere/issues/206)) ([b02c602](https://github.com/validator-labs/validator-plugin-vsphere/commit/b02c602928ab6c1b301e95283cd692edf7a42263))
* **deps:** update azure/setup-helm action to v4 ([#204](https://github.com/validator-labs/validator-plugin-vsphere/issues/204)) ([f7d45a5](https://github.com/validator-labs/validator-plugin-vsphere/commit/f7d45a5996658cf39a2523914d71fd6e4f5e4b2d))
* **deps:** update codecov/codecov-action digest to 0cfda1d ([#200](https://github.com/validator-labs/validator-plugin-vsphere/issues/200)) ([db3d5a0](https://github.com/validator-labs/validator-plugin-vsphere/commit/db3d5a0b9f43b64446e39d6a37219b5337e8bcec))
* **deps:** update codecov/codecov-action digest to 54bcd87 ([#201](https://github.com/validator-labs/validator-plugin-vsphere/issues/201)) ([9c7faa6](https://github.com/validator-labs/validator-plugin-vsphere/commit/9c7faa6423719d23a829e4dd4e401d5b189f6b2e))
* **deps:** update docker/build-push-action digest to af5a7ed ([#208](https://github.com/validator-labs/validator-plugin-vsphere/issues/208)) ([7b904dc](https://github.com/validator-labs/validator-plugin-vsphere/commit/7b904dc980b3eb755e85eb2f9f839c1cbc081659))
* **deps:** update docker/setup-buildx-action digest to 0d103c3 ([#202](https://github.com/validator-labs/validator-plugin-vsphere/issues/202)) ([5411709](https://github.com/validator-labs/validator-plugin-vsphere/commit/54117098215c7acb3461763b9597a1ce500cf63a))
* **deps:** update helm/kind-action action to v1.9.0 ([#190](https://github.com/validator-labs/validator-plugin-vsphere/issues/190)) ([af2f5b4](https://github.com/validator-labs/validator-plugin-vsphere/commit/af2f5b40c36ce1a5c77fdc7680362d72f0fb53ff))
* **deps:** update softprops/action-gh-release action to v2 ([#209](https://github.com/validator-labs/validator-plugin-vsphere/issues/209)) ([4352fcb](https://github.com/validator-labs/validator-plugin-vsphere/commit/4352fcb9e96250400beb9eabbe0b31dfcbe7d73c))
* **deps:** update softprops/action-gh-release digest to d99959e ([#211](https://github.com/validator-labs/validator-plugin-vsphere/issues/211)) ([aeceb3d](https://github.com/validator-labs/validator-plugin-vsphere/commit/aeceb3dd385f47a6e65c2b991cd4dca8727131d4))
* fix broken build link in README ([#203](https://github.com/validator-labs/validator-plugin-vsphere/issues/203)) ([150a118](https://github.com/validator-labs/validator-plugin-vsphere/commit/150a1187b3e92d18b7698fb82fbfc6641e3f4d5f))
* upgrade to validator v0.0.36 ([#210](https://github.com/validator-labs/validator-plugin-vsphere/issues/210)) ([c2512ec](https://github.com/validator-labs/validator-plugin-vsphere/commit/c2512eca74c36aaf567628a952a0195e6bb27957))

## [0.0.17](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.16...v0.0.17) (2024-02-06)


### Other

* update validator ([c6f847c](https://github.com/validator-labs/validator-plugin-vsphere/commit/c6f847ca92bb48b2b4dd71385a5a0a565885d9d2))

## [0.0.16](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.15...v0.0.16) (2024-02-05)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 02704c9 ([#148](https://github.com/validator-labs/validator-plugin-vsphere/issues/148)) ([ecf0ca2](https://github.com/validator-labs/validator-plugin-vsphere/commit/ecf0ca25a17122f5995656974faf9cf150f59533))
* **deps:** update golang.org/x/exp digest to 0dcbfd6 ([#157](https://github.com/validator-labs/validator-plugin-vsphere/issues/157)) ([149f563](https://github.com/validator-labs/validator-plugin-vsphere/commit/149f56302978033129420d36ca596d86bde16ae9))
* **deps:** update golang.org/x/exp digest to 1b97071 ([#169](https://github.com/validator-labs/validator-plugin-vsphere/issues/169)) ([d1f79d4](https://github.com/validator-labs/validator-plugin-vsphere/commit/d1f79d40eecd59d46c3fe7862be9b297b7319704))
* **deps:** update golang.org/x/exp digest to aacd6d4 ([#135](https://github.com/validator-labs/validator-plugin-vsphere/issues/135)) ([3bb6bba](https://github.com/validator-labs/validator-plugin-vsphere/commit/3bb6bbac419179501e946542581c78fc9764c282))
* **deps:** update golang.org/x/exp digest to be819d1 ([#152](https://github.com/validator-labs/validator-plugin-vsphere/issues/152)) ([0d81cb6](https://github.com/validator-labs/validator-plugin-vsphere/commit/0d81cb68e4c0320cf65e99e18dfce4eb7e8a72dd))
* **deps:** update golang.org/x/exp digest to db7319d ([#161](https://github.com/validator-labs/validator-plugin-vsphere/issues/161)) ([3468e79](https://github.com/validator-labs/validator-plugin-vsphere/commit/3468e792f42f7105b848a23ff4da1a40b1acfabb))
* **deps:** update golang.org/x/exp digest to dc181d7 ([#139](https://github.com/validator-labs/validator-plugin-vsphere/issues/139)) ([85e9d4e](https://github.com/validator-labs/validator-plugin-vsphere/commit/85e9d4e67c29242dca909e578a9f1be66f48c095))
* **deps:** update golang.org/x/exp digest to f3f8817 ([#128](https://github.com/validator-labs/validator-plugin-vsphere/issues/128)) ([9c22411](https://github.com/validator-labs/validator-plugin-vsphere/commit/9c224114da63c4b25f7e25c538f1e08a37794ad3))
* **deps:** update module github.com/go-logr/logr to v1.4.0 ([#146](https://github.com/validator-labs/validator-plugin-vsphere/issues/146)) ([055785e](https://github.com/validator-labs/validator-plugin-vsphere/commit/055785e627df6e08b4e6ec6325b15f6da319aba0))
* **deps:** update module github.com/go-logr/logr to v1.4.1 ([#147](https://github.com/validator-labs/validator-plugin-vsphere/issues/147)) ([114938f](https://github.com/validator-labs/validator-plugin-vsphere/commit/114938f213ea7cad5580e41a3dc4cc2698a68e42))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.14.0 ([#160](https://github.com/validator-labs/validator-plugin-vsphere/issues/160)) ([7910806](https://github.com/validator-labs/validator-plugin-vsphere/commit/79108065c045f0ab1eeb6bee897be18eb146b144))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.15.0 ([#165](https://github.com/validator-labs/validator-plugin-vsphere/issues/165)) ([87567cf](https://github.com/validator-labs/validator-plugin-vsphere/commit/87567cf9a55589bbfa0fd4fa1bc413435d8d0288))
* **deps:** update module github.com/onsi/gomega to v1.31.0 ([#166](https://github.com/validator-labs/validator-plugin-vsphere/issues/166)) ([f1a50e6](https://github.com/validator-labs/validator-plugin-vsphere/commit/f1a50e6d1edd4a68adeac4229802c80f9b4abcdf))
* **deps:** update module github.com/onsi/gomega to v1.31.1 ([#170](https://github.com/validator-labs/validator-plugin-vsphere/issues/170)) ([b19c2bf](https://github.com/validator-labs/validator-plugin-vsphere/commit/b19c2bf66b3e397240db47fa7fb19122d0be14af))
* **deps:** update module github.com/validator-labs/validator to v0.0.28 ([#120](https://github.com/validator-labs/validator-plugin-vsphere/issues/120)) ([e5f0da5](https://github.com/validator-labs/validator-plugin-vsphere/commit/e5f0da5c8b7733c3065568302c20832602a21d5e))
* **deps:** update module github.com/validator-labs/validator to v0.0.30 ([#140](https://github.com/validator-labs/validator-plugin-vsphere/issues/140)) ([d6eb777](https://github.com/validator-labs/validator-plugin-vsphere/commit/d6eb777ee91f2accfa662e087e8c10583fdb8392))
* **deps:** update module github.com/validator-labs/validator to v0.0.31 ([#149](https://github.com/validator-labs/validator-plugin-vsphere/issues/149)) ([7a3335d](https://github.com/validator-labs/validator-plugin-vsphere/commit/7a3335d109f3d269a272395a26344d941b145c92))
* **deps:** update module github.com/validator-labs/validator to v0.0.32 ([#150](https://github.com/validator-labs/validator-plugin-vsphere/issues/150)) ([d20321b](https://github.com/validator-labs/validator-plugin-vsphere/commit/d20321bd1a53a9d3702154b4524b3f05f66d9b73))
* **deps:** update module github.com/vmware/govmomi to v0.34.0 ([#133](https://github.com/validator-labs/validator-plugin-vsphere/issues/133)) ([0bcfc6f](https://github.com/validator-labs/validator-plugin-vsphere/commit/0bcfc6f7062a157931f8f1b7f1a4f3004b7928ea))
* **deps:** update module github.com/vmware/govmomi to v0.34.1 ([#142](https://github.com/validator-labs/validator-plugin-vsphere/issues/142)) ([04c8a18](https://github.com/validator-labs/validator-plugin-vsphere/commit/04c8a18a2c870d4ee00d0dce0a2d361b107663bb))
* **deps:** update module github.com/vmware/govmomi to v0.34.2 ([#153](https://github.com/validator-labs/validator-plugin-vsphere/issues/153)) ([a338a82](https://github.com/validator-labs/validator-plugin-vsphere/commit/a338a825fb7b4aef7cf74a71bff8d7ea50e4fc3f))
* update VR's ExpectedResults if/when rules are added to an existing VsphereValidator ([#185](https://github.com/validator-labs/validator-plugin-vsphere/issues/185)) ([a510966](https://github.com/validator-labs/validator-plugin-vsphere/commit/a5109669d23ad9e27765e66549187a700d1abaf7))


### Other

* **deps:** update actions/setup-go action to v5 ([#127](https://github.com/validator-labs/validator-plugin-vsphere/issues/127)) ([ce18cbd](https://github.com/validator-labs/validator-plugin-vsphere/commit/ce18cbd60a780a766dfaa389cb7a463257e6ce08))
* **deps:** update actions/setup-python action to v5 ([#126](https://github.com/validator-labs/validator-plugin-vsphere/issues/126)) ([4381f33](https://github.com/validator-labs/validator-plugin-vsphere/commit/4381f332c7c3ed89dc3aaaa592d253b299ad6323))
* **deps:** update actions/upload-artifact action to v4 ([#136](https://github.com/validator-labs/validator-plugin-vsphere/issues/136)) ([91758cf](https://github.com/validator-labs/validator-plugin-vsphere/commit/91758cfe69c4e2ad1e0484b9f4f702e308f182b3))
* **deps:** update actions/upload-artifact digest to 1eb3cb2 ([#162](https://github.com/validator-labs/validator-plugin-vsphere/issues/162)) ([b68863a](https://github.com/validator-labs/validator-plugin-vsphere/commit/b68863aecdc5f6e74c97672b931c155ff7baed7f))
* **deps:** update actions/upload-artifact digest to 26f96df ([#172](https://github.com/validator-labs/validator-plugin-vsphere/issues/172)) ([2f58c5a](https://github.com/validator-labs/validator-plugin-vsphere/commit/2f58c5af09a630c6555859ad03e4297065debec1))
* **deps:** update actions/upload-artifact digest to 694cdab ([#168](https://github.com/validator-labs/validator-plugin-vsphere/issues/168)) ([d22011d](https://github.com/validator-labs/validator-plugin-vsphere/commit/d22011d59f26bf95a02ba23a7ecd5843b8dfcd85))
* **deps:** update anchore/sbom-action action to v0.15.1 ([#124](https://github.com/validator-labs/validator-plugin-vsphere/issues/124)) ([144d6c3](https://github.com/validator-labs/validator-plugin-vsphere/commit/144d6c3dea7b951449a3e6eab588c02ac093f8b7))
* **deps:** update anchore/sbom-action action to v0.15.2 ([#151](https://github.com/validator-labs/validator-plugin-vsphere/issues/151)) ([988fdde](https://github.com/validator-labs/validator-plugin-vsphere/commit/988fddec9de4755c4e3de4c5cb6fe54478e2b0d3))
* **deps:** update anchore/sbom-action action to v0.15.3 ([#154](https://github.com/validator-labs/validator-plugin-vsphere/issues/154)) ([f13a094](https://github.com/validator-labs/validator-plugin-vsphere/commit/f13a094aba1f87d6a2fa911deba4840d16cd3533))
* **deps:** update anchore/sbom-action action to v0.15.4 ([#167](https://github.com/validator-labs/validator-plugin-vsphere/issues/167)) ([10ac377](https://github.com/validator-labs/validator-plugin-vsphere/commit/10ac3771340239a53cd67e6bfa664b39091c5727))
* **deps:** update anchore/sbom-action action to v0.15.5 ([#171](https://github.com/validator-labs/validator-plugin-vsphere/issues/171)) ([01ed816](https://github.com/validator-labs/validator-plugin-vsphere/commit/01ed81610a9837e9b2770b7d28283e4c4205e443))
* **deps:** update anchore/sbom-action action to v0.15.7 ([#175](https://github.com/validator-labs/validator-plugin-vsphere/issues/175)) ([9b6d96d](https://github.com/validator-labs/validator-plugin-vsphere/commit/9b6d96d556e25d010b8b4e55606b7fce9d8c9c6b))
* **deps:** update anchore/sbom-action action to v0.15.8 ([#177](https://github.com/validator-labs/validator-plugin-vsphere/issues/177)) ([900e59c](https://github.com/validator-labs/validator-plugin-vsphere/commit/900e59c2540b353a07202b84d19e2f7773c935ca))
* **deps:** update codecov/codecov-action digest to 4fe8c5f ([#174](https://github.com/validator-labs/validator-plugin-vsphere/issues/174)) ([f091db3](https://github.com/validator-labs/validator-plugin-vsphere/commit/f091db3e192472065eb8f0ee91d4e1f759b8ce3f))
* **deps:** update codecov/codecov-action digest to ab904c4 ([#176](https://github.com/validator-labs/validator-plugin-vsphere/issues/176)) ([6c2df82](https://github.com/validator-labs/validator-plugin-vsphere/commit/6c2df82245ce5a2e391654743406807058859ecd))
* **deps:** update codecov/codecov-action digest to e0b68c6 ([#183](https://github.com/validator-labs/validator-plugin-vsphere/issues/183)) ([8d443f6](https://github.com/validator-labs/validator-plugin-vsphere/commit/8d443f6d7b3517eb9c07ef05e37cae502ad26470))
* **deps:** update gcr.io/spectro-images-public/golang docker tag to v1.22 ([#156](https://github.com/validator-labs/validator-plugin-vsphere/issues/156)) ([673cdf0](https://github.com/validator-labs/validator-plugin-vsphere/commit/673cdf07e329e2343b72b1746e49589cf314240d))
* **deps:** update golang:1.21 docker digest to 11dc65e ([#141](https://github.com/validator-labs/validator-plugin-vsphere/issues/141)) ([06faef9](https://github.com/validator-labs/validator-plugin-vsphere/commit/06faef93c6ea316e03dc189225b097ab3e7a5bdf))
* **deps:** update golang:1.21 docker digest to 1a9d253 ([#143](https://github.com/validator-labs/validator-plugin-vsphere/issues/143)) ([712eb58](https://github.com/validator-labs/validator-plugin-vsphere/commit/712eb58f8e5b64b2b7a3b389dc9883ec9e5939ac))
* **deps:** update golang:1.21 docker digest to 21260a4 ([#158](https://github.com/validator-labs/validator-plugin-vsphere/issues/158)) ([df280c4](https://github.com/validator-labs/validator-plugin-vsphere/commit/df280c46ce4a7843c8e38b3b7be1a76f7aeb41be))
* **deps:** update golang:1.21 docker digest to 25b05d5 ([#138](https://github.com/validator-labs/validator-plugin-vsphere/issues/138)) ([8031596](https://github.com/validator-labs/validator-plugin-vsphere/commit/8031596671d6ddc1e41a13e7773827e3a8e48945))
* **deps:** update golang:1.21 docker digest to 2ff79bc ([#132](https://github.com/validator-labs/validator-plugin-vsphere/issues/132)) ([80e09ca](https://github.com/validator-labs/validator-plugin-vsphere/commit/80e09ca7620eb4481ebe3f86c4f3b9143c859f6a))
* **deps:** update golang:1.21 docker digest to 3c871ba ([#180](https://github.com/validator-labs/validator-plugin-vsphere/issues/180)) ([afeede0](https://github.com/validator-labs/validator-plugin-vsphere/commit/afeede0d724b1f3a287e2da2eb0b3a8d93e7f2d9))
* **deps:** update golang:1.21 docker digest to 4d1942c ([#181](https://github.com/validator-labs/validator-plugin-vsphere/issues/181)) ([ba7a426](https://github.com/validator-labs/validator-plugin-vsphere/commit/ba7a4267b8085f83142827577cbd3ab41476782e))
* **deps:** update golang:1.21 docker digest to 58e14a9 ([#125](https://github.com/validator-labs/validator-plugin-vsphere/issues/125)) ([6678369](https://github.com/validator-labs/validator-plugin-vsphere/commit/66783694001853a027dae6ecd6401ffaabfbda34))
* **deps:** update golang:1.21 docker digest to 5f5d61d ([#164](https://github.com/validator-labs/validator-plugin-vsphere/issues/164)) ([edf8dd9](https://github.com/validator-labs/validator-plugin-vsphere/commit/edf8dd9fb5e30735dd4928c439960234d2762e18))
* **deps:** update golang:1.21 docker digest to 672a228 ([#145](https://github.com/validator-labs/validator-plugin-vsphere/issues/145)) ([ca19d48](https://github.com/validator-labs/validator-plugin-vsphere/commit/ca19d482c6f63666a437687324139db738a2aae3))
* **deps:** update golang:1.21 docker digest to 6fbd2d3 ([#159](https://github.com/validator-labs/validator-plugin-vsphere/issues/159)) ([b82d606](https://github.com/validator-labs/validator-plugin-vsphere/commit/b82d606e6483c70a4381317dd67fc703234083c7))
* **deps:** update golang:1.21 docker digest to 7026fb7 ([#155](https://github.com/validator-labs/validator-plugin-vsphere/issues/155)) ([9275b49](https://github.com/validator-labs/validator-plugin-vsphere/commit/9275b496c5761c850a52f5ad64f5024cb7dcb5ad))
* **deps:** update golang:1.21 docker digest to 76aadd9 ([#173](https://github.com/validator-labs/validator-plugin-vsphere/issues/173)) ([30cf9b9](https://github.com/validator-labs/validator-plugin-vsphere/commit/30cf9b9e69adabb27f0c9acdad12efda97931859))
* **deps:** update golang:1.21 docker digest to 7b575fe ([#182](https://github.com/validator-labs/validator-plugin-vsphere/issues/182)) ([4482055](https://github.com/validator-labs/validator-plugin-vsphere/commit/4482055b20561858bc9fa503d8202225543e8355))
* **deps:** update golang:1.21 docker digest to ae34fbf ([#130](https://github.com/validator-labs/validator-plugin-vsphere/issues/130)) ([3ea9b29](https://github.com/validator-labs/validator-plugin-vsphere/commit/3ea9b29f29d4cda466b5d85804cb00875f7b6096))
* **deps:** update golang:1.21 docker digest to fb02af5 ([#144](https://github.com/validator-labs/validator-plugin-vsphere/issues/144)) ([e0af841](https://github.com/validator-labs/validator-plugin-vsphere/commit/e0af841383283aea7ee3f8ff9beed003faa4608f))
* **deps:** update google-github-actions/release-please-action action to v4 ([#122](https://github.com/validator-labs/validator-plugin-vsphere/issues/122)) ([2061c1f](https://github.com/validator-labs/validator-plugin-vsphere/commit/2061c1f5e87aee29e2e386eace87971d6c172d82))
* **deps:** update google-github-actions/release-please-action digest to a2d8d68 ([#129](https://github.com/validator-labs/validator-plugin-vsphere/issues/129)) ([de9f611](https://github.com/validator-labs/validator-plugin-vsphere/commit/de9f611da73dd1561b2afd7c6a560f6e82e80c61))
* **deps:** update google-github-actions/release-please-action digest to cc61a07 ([#137](https://github.com/validator-labs/validator-plugin-vsphere/issues/137)) ([0b719ce](https://github.com/validator-labs/validator-plugin-vsphere/commit/0b719cef766e7b368d1901db155106db18436a9f))

## [0.0.15](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.14...v0.0.15) (2023-11-30)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 6522937 ([#115](https://github.com/validator-labs/validator-plugin-vsphere/issues/115)) ([5a72d63](https://github.com/validator-labs/validator-plugin-vsphere/commit/5a72d6333c8e9e701019a75b46d8ef01b8653b7a))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.2 ([#116](https://github.com/validator-labs/validator-plugin-vsphere/issues/116)) ([252f7fc](https://github.com/validator-labs/validator-plugin-vsphere/commit/252f7fcdc55d519f413658566bc78a4e8910ec02))


### Other

* **deps:** update actions/checkout action to v4 ([#119](https://github.com/validator-labs/validator-plugin-vsphere/issues/119)) ([7c612c4](https://github.com/validator-labs/validator-plugin-vsphere/commit/7c612c45da18a2f85e3c35d9fb0aee07bc54df90))
* **deps:** update anchore/sbom-action action to v0.15.0 ([#113](https://github.com/validator-labs/validator-plugin-vsphere/issues/113)) ([ad9df0d](https://github.com/validator-labs/validator-plugin-vsphere/commit/ad9df0d7d7e2afbc3a4954cedfba58d4a1ed21ae))
* **deps:** update golang:1.21 docker digest to 9baee0e ([#114](https://github.com/validator-labs/validator-plugin-vsphere/issues/114)) ([b3ce98f](https://github.com/validator-labs/validator-plugin-vsphere/commit/b3ce98f989859e90570dfb67395b56ae262243ca))
* Update validator core and refactor ([#121](https://github.com/validator-labs/validator-plugin-vsphere/issues/121)) ([6d0d6f8](https://github.com/validator-labs/validator-plugin-vsphere/commit/6d0d6f888d0e91963294fc39ca35288fc9de4792))

## [0.0.14](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.13...v0.0.14) (2023-11-17)


### Refactoring

* remove auth secret creation ([#109](https://github.com/validator-labs/validator-plugin-vsphere/issues/109)) ([965ed7f](https://github.com/validator-labs/validator-plugin-vsphere/commit/965ed7f63874df4ac5280edecbe41df6d9d1999f))

## [0.0.13](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.12...v0.0.13) (2023-11-17)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 9a3e603 ([#102](https://github.com/validator-labs/validator-plugin-vsphere/issues/102)) ([9456e5b](https://github.com/validator-labs/validator-plugin-vsphere/commit/9456e5b8664a5b8df18a1502a21748cb6f49ae8e))
* **deps:** update kubernetes packages to v0.28.4 ([#106](https://github.com/validator-labs/validator-plugin-vsphere/issues/106)) ([2a9fd27](https://github.com/validator-labs/validator-plugin-vsphere/commit/2a9fd2786aa58527acecf4264a998c605ea70634))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.1 ([#105](https://github.com/validator-labs/validator-plugin-vsphere/issues/105)) ([972f639](https://github.com/validator-labs/validator-plugin-vsphere/commit/972f639a35077137806267befe9afde5b0f7b301))
* **deps:** update module github.com/onsi/gomega to v1.30.0 ([#101](https://github.com/validator-labs/validator-plugin-vsphere/issues/101)) ([99d4cca](https://github.com/validator-labs/validator-plugin-vsphere/commit/99d4ccac73fc906fb45edb0dd1b1ee4190d4b699))
* **deps:** update module github.com/vmware/govmomi to v0.33.1 ([#92](https://github.com/validator-labs/validator-plugin-vsphere/issues/92)) ([70f7980](https://github.com/validator-labs/validator-plugin-vsphere/commit/70f79803e7ebdbb8474c33697eb1896d9ed6a909))
* update VsphereValidator CRD ([0dccf7d](https://github.com/validator-labs/validator-plugin-vsphere/commit/0dccf7dd761e5bf60f33e8b00e7a123bf23cccf6))


### Other

* **deps:** update docker/build-push-action digest to 4a13e50 ([#108](https://github.com/validator-labs/validator-plugin-vsphere/issues/108)) ([6a5abdb](https://github.com/validator-labs/validator-plugin-vsphere/commit/6a5abdb51c9a9119a65cf685f3e065041818415a))
* **deps:** update golang:1.21 docker digest to 5206873 ([#96](https://github.com/validator-labs/validator-plugin-vsphere/issues/96)) ([1d90fbb](https://github.com/validator-labs/validator-plugin-vsphere/commit/1d90fbb28446890b630028cd958bad04a9d559fb))
* **deps:** update golang:1.21 docker digest to 57bf74a ([#107](https://github.com/validator-labs/validator-plugin-vsphere/issues/107)) ([9dd1218](https://github.com/validator-labs/validator-plugin-vsphere/commit/9dd121891443c988f1b2d18ade76817eeb0a7e0a))
* **deps:** update golang:1.21 docker digest to 81cd210 ([#100](https://github.com/validator-labs/validator-plugin-vsphere/issues/100)) ([d6d4fd4](https://github.com/validator-labs/validator-plugin-vsphere/commit/d6d4fd43581f2f9360af32195072b2c518a18958))
* **deps:** update golang:1.21 docker digest to 84e41b3 ([#95](https://github.com/validator-labs/validator-plugin-vsphere/issues/95)) ([37aaee1](https://github.com/validator-labs/validator-plugin-vsphere/commit/37aaee191731a459f497d1ca5f9481ad624f4d74))
* **deps:** update golang:1.21 docker digest to b113af1 ([#97](https://github.com/validator-labs/validator-plugin-vsphere/issues/97)) ([7a03507](https://github.com/validator-labs/validator-plugin-vsphere/commit/7a035072501d3b034194ba9d7f66046769dbb271))
* **deps:** update google-github-actions/release-please-action digest to db8f2c6 ([#99](https://github.com/validator-labs/validator-plugin-vsphere/issues/99)) ([4e61952](https://github.com/validator-labs/validator-plugin-vsphere/commit/4e61952cb09c6d4dbbef52e8b869fdedfe268880))
* **deps:** update helm/chart-testing-action action to v2.6.0 ([#94](https://github.com/validator-labs/validator-plugin-vsphere/issues/94)) ([e8addc3](https://github.com/validator-labs/validator-plugin-vsphere/commit/e8addc3e585bed76368b1359b341f2625e3e40ed))
* **deps:** update helm/chart-testing-action action to v2.6.1 ([#98](https://github.com/validator-labs/validator-plugin-vsphere/issues/98)) ([0f24db3](https://github.com/validator-labs/validator-plugin-vsphere/commit/0f24db3b089867089380ee473221e12d7a242305))

## [0.0.12](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.11...v0.0.12) (2023-10-27)


### Features

* Add NTP validation to validator-plugin-vsphere ([#88](https://github.com/validator-labs/validator-plugin-vsphere/issues/88)) ([b439ebb](https://github.com/validator-labs/validator-plugin-vsphere/commit/b439ebb8585947a798f036a0803960520daf1f24))


### Bug Fixes

* **deps:** update module github.com/go-logr/logr to v1.3.0 ([#91](https://github.com/validator-labs/validator-plugin-vsphere/issues/91)) ([a84f798](https://github.com/validator-labs/validator-plugin-vsphere/commit/a84f7985637c3f0ab28e35b9d339dbfbe8793ccf))
* **deps:** update module github.com/onsi/gomega to v1.28.1 ([#87](https://github.com/validator-labs/validator-plugin-vsphere/issues/87)) ([e264ce9](https://github.com/validator-labs/validator-plugin-vsphere/commit/e264ce9d2ac047a1fdb2f3621e2e70139e2b202c))
* **deps:** update module github.com/onsi/gomega to v1.29.0 ([#89](https://github.com/validator-labs/validator-plugin-vsphere/issues/89)) ([51f0109](https://github.com/validator-labs/validator-plugin-vsphere/commit/51f010975192135478794f991ec94de8c935a633))
* **deps:** update module github.com/vmware/govmomi to v0.33.0 ([#90](https://github.com/validator-labs/validator-plugin-vsphere/issues/90)) ([7e25a95](https://github.com/validator-labs/validator-plugin-vsphere/commit/7e25a95cef4fdb6a0aff3675e3eb353b188d4e02))


### Other

* add license ([8fb9e2e](https://github.com/validator-labs/validator-plugin-vsphere/commit/8fb9e2e537221cb260c22f23e4bfd81ead41ba18))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.15.0 ([#85](https://github.com/validator-labs/validator-plugin-vsphere/issues/85)) ([33f9f8c](https://github.com/validator-labs/validator-plugin-vsphere/commit/33f9f8c506f37555add0158a64050e5921dddae4))
* update plugin code ([8c29d0c](https://github.com/validator-labs/validator-plugin-vsphere/commit/8c29d0c93db6e71063b7a3ec6277dc8098a58c40))

## [0.0.11](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.10...v0.0.11) (2023-10-20)


### Bug Fixes

* ct lints ([739b4a8](https://github.com/validator-labs/validator-plugin-vsphere/commit/739b4a80e8b0ba0723213e09b94ba1b7fd97ea2f))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.3 ([#82](https://github.com/validator-labs/validator-plugin-vsphere/issues/82)) ([f031533](https://github.com/validator-labs/validator-plugin-vsphere/commit/f03153304f8c9a331242b255178749aef8b1fe48))


### Other

* **deps:** bump golang.org/x/net from 0.16.0 to 0.17.0 ([#74](https://github.com/validator-labs/validator-plugin-vsphere/issues/74)) ([a27698d](https://github.com/validator-labs/validator-plugin-vsphere/commit/a27698d6a0df19c892880dac7f83512f91d24e0e))
* **deps:** update actions/checkout digest to b4ffde6 ([#80](https://github.com/validator-labs/validator-plugin-vsphere/issues/80)) ([0b5d05d](https://github.com/validator-labs/validator-plugin-vsphere/commit/0b5d05d24901f870e66ac30d20a70216c3f99722))
* **deps:** update actions/setup-python digest to 65d7f2d ([#78](https://github.com/validator-labs/validator-plugin-vsphere/issues/78)) ([5652f1d](https://github.com/validator-labs/validator-plugin-vsphere/commit/5652f1d641f3f53253858703f5be857b35ac9dc9))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.14.4 ([#71](https://github.com/validator-labs/validator-plugin-vsphere/issues/71)) ([db4676e](https://github.com/validator-labs/validator-plugin-vsphere/commit/db4676e5d4621da7910c1cbefda5f1ac414b0d5f))
* **deps:** update google-github-actions/release-please-action digest to 4c5670f ([#79](https://github.com/validator-labs/validator-plugin-vsphere/issues/79)) ([93fb4a8](https://github.com/validator-labs/validator-plugin-vsphere/commit/93fb4a8659650db5a45353830b1985dd1baae286))
* enable renovate automerges ([2639694](https://github.com/validator-labs/validator-plugin-vsphere/commit/2639694fac061f5e284f1fccc48a675b0738fdef))
* release 0.0.11 ([dffec26](https://github.com/validator-labs/validator-plugin-vsphere/commit/dffec2605f8cf5db33522a6c7ff772fd48e37193))


### Refactoring

* validator -&gt; validator ([#83](https://github.com/validator-labs/validator-plugin-vsphere/issues/83)) ([acf1f53](https://github.com/validator-labs/validator-plugin-vsphere/commit/acf1f53d94f209fd22da31fec62f34b5afee6b53))

## [0.0.10](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.9...v0.0.10) (2023-10-16)


### Bug Fixes

* Fix dockerfile for refactor of vsphere from internal to pkg ([#70](https://github.com/validator-labs/validator-plugin-vsphere/issues/70)) ([a08e02e](https://github.com/validator-labs/validator-plugin-vsphere/commit/a08e02eddfc428e47cd271cdd4b06408b3aff73c))
* move vsphere libs to pkg/ so they can be used by other projects ([#69](https://github.com/validator-labs/validator-plugin-vsphere/issues/69)) ([1bd8012](https://github.com/validator-labs/validator-plugin-vsphere/commit/1bd801235804e5100291b69cb7807d1b4b066c84))


### Other

* **deps:** update golang:1.21 docker digest to 02d7116 ([#67](https://github.com/validator-labs/validator-plugin-vsphere/issues/67)) ([d0a82c6](https://github.com/validator-labs/validator-plugin-vsphere/commit/d0a82c6d55c56b8a6b2bf04e01a78fe4d04cbe38))
* **deps:** update golang:1.21 docker digest to 24a0937 ([#68](https://github.com/validator-labs/validator-plugin-vsphere/issues/68)) ([5bd2734](https://github.com/validator-labs/validator-plugin-vsphere/commit/5bd27348f44f56c6daaebcc9ed7b11aed5374708))
* **deps:** update golang:1.21 docker digest to 4d5cf6c ([#65](https://github.com/validator-labs/validator-plugin-vsphere/issues/65)) ([0dacc3b](https://github.com/validator-labs/validator-plugin-vsphere/commit/0dacc3bf964b3e408c6a4a3ed4a989ad7442555e))

## [0.0.9](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.8...v0.0.9) (2023-10-10)


### Bug Fixes

* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.0 ([#61](https://github.com/validator-labs/validator-plugin-vsphere/issues/61)) ([f927c76](https://github.com/validator-labs/validator-plugin-vsphere/commit/f927c7660e8238a51f48abacf845e63180c3c2cb))
* **deps:** update module github.com/validator-labs/validator to v0.0.9 ([#64](https://github.com/validator-labs/validator-plugin-vsphere/issues/64)) ([fdb3027](https://github.com/validator-labs/validator-plugin-vsphere/commit/fdb30270d936f4cfa29d4bfaecd2d41e1356910c))


### Other

* better log messages in controller ([#60](https://github.com/validator-labs/validator-plugin-vsphere/issues/60)) ([2e64ce3](https://github.com/validator-labs/validator-plugin-vsphere/commit/2e64ce3dd1e3fd74e6e89bfc127ad72741b5e4a3))
* **deps:** update golang:1.21 docker digest to e9ebfe9 ([#42](https://github.com/validator-labs/validator-plugin-vsphere/issues/42)) ([204b045](https://github.com/validator-labs/validator-plugin-vsphere/commit/204b045a018fb93359bbbf099b3df7acc26db785))

## [0.0.8](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.7...v0.0.8) (2023-10-09)


### Bug Fixes

* release please comments in chart.yaml and values.yaml ([#58](https://github.com/validator-labs/validator-plugin-vsphere/issues/58)) ([6b0ad05](https://github.com/validator-labs/validator-plugin-vsphere/commit/6b0ad0550b6753721030ddc6a13ce119fa9ed2c3))

## [0.0.7](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.6...v0.0.7) (2023-10-09)


### Bug Fixes

* update charts and add proper templating for auth secret ([#55](https://github.com/validator-labs/validator-plugin-vsphere/issues/55)) ([3ba9c4b](https://github.com/validator-labs/validator-plugin-vsphere/commit/3ba9c4b2e4b9a1a00e659c81da4185f837b814fc))

## [0.0.6](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.5...v0.0.6) (2023-10-06)


### Bug Fixes

* yaml tag for auth secretName ([#53](https://github.com/validator-labs/validator-plugin-vsphere/issues/53)) ([cec752f](https://github.com/validator-labs/validator-plugin-vsphere/commit/cec752fa55f23748c5943a32d065142a4f41fabf))

## [0.0.5](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.4...v0.0.5) (2023-10-06)


### Bug Fixes

* fix yaml tag for nodepool cpu ([e21d8dc](https://github.com/validator-labs/validator-plugin-vsphere/commit/e21d8dcf04098d428732b348a5bf22f27092330e))


### Other

* release 0.0.4 ([#50](https://github.com/validator-labs/validator-plugin-vsphere/issues/50)) ([c419c9d](https://github.com/validator-labs/validator-plugin-vsphere/commit/c419c9d4e9298ee8127ad884c2d70d00aa3b5b87))
* release 0.0.5 ([#51](https://github.com/validator-labs/validator-plugin-vsphere/issues/51)) ([04508a8](https://github.com/validator-labs/validator-plugin-vsphere/commit/04508a88d66c6ea42aaa9162fc9e5c939dba7cf2))

## [0.0.4](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.3...v0.0.4) (2023-10-06)


### Bug Fixes

* fix generated code and temporarily disable tests ([#46](https://github.com/validator-labs/validator-plugin-vsphere/issues/46)) ([56cf9a7](https://github.com/validator-labs/validator-plugin-vsphere/commit/56cf9a715086f30fd952d98c449cd8df31dae6c0))


### Other

* Disable roleprivilege tests temporarily ([#48](https://github.com/validator-labs/validator-plugin-vsphere/issues/48)) ([3d4d736](https://github.com/validator-labs/validator-plugin-vsphere/commit/3d4d73622a6c0ab46b7cd288ed76a2558ad21bf9))

## [0.0.3](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.2...v0.0.3) (2023-10-06)


### Features

* Add support to validate arbitrary user's role and entity privileges instead of the one specified under auth secret ([#41](https://github.com/validator-labs/validator-plugin-vsphere/issues/41)) ([033f665](https://github.com/validator-labs/validator-plugin-vsphere/commit/033f665794dfadbd4d1473c7fdaed1242d7d0669))


### Other

* Add yaml tags to api types ([#44](https://github.com/validator-labs/validator-plugin-vsphere/issues/44)) ([1578a1f](https://github.com/validator-labs/validator-plugin-vsphere/commit/1578a1f43992f7fa25ce0316431dc39c5e18d5ad))

## [0.0.2](https://github.com/validator-labs/validator-plugin-vsphere/compare/v0.0.1...v0.0.2) (2023-10-02)


### Features

* Add support to validate available resources on cluster, hosts and resourcepools ([#26](https://github.com/validator-labs/validator-plugin-vsphere/issues/26)) ([d15e5a4](https://github.com/validator-labs/validator-plugin-vsphere/commit/d15e5a4a3ce7fc1bbe898dacff6f53388a9356ae))


### Bug Fixes

* better logging and missing fields in cmd/main.go ([#10](https://github.com/validator-labs/validator-plugin-vsphere/issues/10)) ([bb39b0b](https://github.com/validator-labs/validator-plugin-vsphere/commit/bb39b0b0a4d12cc6554041f86442e9115ba93889))
* **deps:** update kubernetes packages to v0.28.2 ([#28](https://github.com/validator-labs/validator-plugin-vsphere/issues/28)) ([cd84314](https://github.com/validator-labs/validator-plugin-vsphere/commit/cd84314cec33ac51d2f7a9f75ca851edfa50359b))
* **deps:** update module github.com/onsi/ginkgo to v2 ([#39](https://github.com/validator-labs/validator-plugin-vsphere/issues/39)) ([0251709](https://github.com/validator-labs/validator-plugin-vsphere/commit/025170979179cd839cf967a71cce29ee00961a61))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.1 ([#33](https://github.com/validator-labs/validator-plugin-vsphere/issues/33)) ([4303aed](https://github.com/validator-labs/validator-plugin-vsphere/commit/4303aed9d4c53c6eb764b39262b464480ee51874))
* **deps:** update module github.com/onsi/gomega to v1.27.10 ([#1](https://github.com/validator-labs/validator-plugin-vsphere/issues/1)) ([f0579a8](https://github.com/validator-labs/validator-plugin-vsphere/commit/f0579a804a165d4b568cb95e997cb315b70cfab5))
* **deps:** update module github.com/onsi/gomega to v1.28.0 ([#36](https://github.com/validator-labs/validator-plugin-vsphere/issues/36)) ([14b3f34](https://github.com/validator-labs/validator-plugin-vsphere/commit/14b3f3477f59ddd1684f088b79dee8ab12602347))
* **deps:** update module github.com/sirupsen/logrus to v1.9.3 ([#34](https://github.com/validator-labs/validator-plugin-vsphere/issues/34)) ([588e237](https://github.com/validator-labs/validator-plugin-vsphere/commit/588e2370111567e3548c038d098bbe7bfebf8cbd))
* **deps:** update module github.com/validator-labs/validator to v0.0.6 ([#7](https://github.com/validator-labs/validator-plugin-vsphere/issues/7)) ([ff931ed](https://github.com/validator-labs/validator-plugin-vsphere/commit/ff931edd2782e664149a6c51c67e4d2364489ef3))
* **deps:** update module github.com/validator-labs/validator to v0.0.8 ([#20](https://github.com/validator-labs/validator-plugin-vsphere/issues/20)) ([9c54342](https://github.com/validator-labs/validator-plugin-vsphere/commit/9c54342788a302ea591c630d272fbd7e2471d02a))
* **deps:** update module github.com/vmware/govmomi to v0.31.0 ([#37](https://github.com/validator-labs/validator-plugin-vsphere/issues/37)) ([530cca0](https://github.com/validator-labs/validator-plugin-vsphere/commit/530cca01ba680dff1207b3629a390a42cb33937f))
* **deps:** update module github.com/vmware/govmomi to v0.32.0 ([#40](https://github.com/validator-labs/validator-plugin-vsphere/issues/40)) ([e4d2478](https://github.com/validator-labs/validator-plugin-vsphere/commit/e4d2478e5d3be3fc382b0e197b588dee54a66b56))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.2 ([#35](https://github.com/validator-labs/validator-plugin-vsphere/issues/35)) ([8327ed6](https://github.com/validator-labs/validator-plugin-vsphere/commit/8327ed6ec6446ad5c73f8c1cd24485ec687ea498))
* issues with updating validationresult and chart fixes ([#9](https://github.com/validator-labs/validator-plugin-vsphere/issues/9)) ([6cfbc56](https://github.com/validator-labs/validator-plugin-vsphere/commit/6cfbc569ae551da357593b2bb74a6d8f06838c43))


### Other

* Add charts ([#8](https://github.com/validator-labs/validator-plugin-vsphere/issues/8)) ([a0584bd](https://github.com/validator-labs/validator-plugin-vsphere/commit/a0584bd7e59ca2fadf5f7fd8d706fecfe928d5a5))
* add github workflows ([#11](https://github.com/validator-labs/validator-plugin-vsphere/issues/11)) ([fdc6b8f](https://github.com/validator-labs/validator-plugin-vsphere/commit/fdc6b8fb3f2682f58b52bf23eb2cc6f68aee0c59))
* add pre-commit config ([0e3cfb3](https://github.com/validator-labs/validator-plugin-vsphere/commit/0e3cfb3ed8760e76bdf8d68419d062be0c2d4b9b))
* Add release-please-config ([#29](https://github.com/validator-labs/validator-plugin-vsphere/issues/29)) ([27ca573](https://github.com/validator-labs/validator-plugin-vsphere/commit/27ca573fd3d5e8d526b75dc469b44149192b1c02))
* configure renovate ([6de3e6b](https://github.com/validator-labs/validator-plugin-vsphere/commit/6de3e6b713ca065b47268fe9e9e0c24bec044c51))
* **deps:** update actions/checkout action to v4 ([#17](https://github.com/validator-labs/validator-plugin-vsphere/issues/17)) ([fab282f](https://github.com/validator-labs/validator-plugin-vsphere/commit/fab282fa3d32ad7d1b42ba5417809da001be61b8))
* **deps:** update actions/checkout digest to f43a0e5 ([#14](https://github.com/validator-labs/validator-plugin-vsphere/issues/14)) ([f12efa7](https://github.com/validator-labs/validator-plugin-vsphere/commit/f12efa7108e25e4e524546e12926df33ae55484f))
* **deps:** update actions/setup-go digest to 93397be ([#15](https://github.com/validator-labs/validator-plugin-vsphere/issues/15)) ([3df0d01](https://github.com/validator-labs/validator-plugin-vsphere/commit/3df0d0104b7c4de8cd63e325e2edbf478dc90b22))
* **deps:** update actions/upload-artifact digest to a8a3f3a ([#18](https://github.com/validator-labs/validator-plugin-vsphere/issues/18)) ([1245a73](https://github.com/validator-labs/validator-plugin-vsphere/commit/1245a738ecb1c72e2b95bb60109429be087eddd5))
* **deps:** update docker/build-push-action action to v5 ([#22](https://github.com/validator-labs/validator-plugin-vsphere/issues/22)) ([f258548](https://github.com/validator-labs/validator-plugin-vsphere/commit/f25854898b324b0ef9bef50bd8062494740f054c))
* **deps:** update docker/build-push-action digest to 0a97817 ([#21](https://github.com/validator-labs/validator-plugin-vsphere/issues/21)) ([da721a1](https://github.com/validator-labs/validator-plugin-vsphere/commit/da721a117ae351db07126d2d73f9bcd520d49cc6))
* **deps:** update docker/login-action action to v3 ([#23](https://github.com/validator-labs/validator-plugin-vsphere/issues/23)) ([191a50e](https://github.com/validator-labs/validator-plugin-vsphere/commit/191a50e2d7c53dbe592139f11ec460e060a06362))
* **deps:** update docker/setup-buildx-action action to v3 ([#38](https://github.com/validator-labs/validator-plugin-vsphere/issues/38)) ([97446de](https://github.com/validator-labs/validator-plugin-vsphere/commit/97446dea7e18d4ca43e127ad17353ca5bfa38867))
* **deps:** update docker/setup-buildx-action digest to 885d146 ([#16](https://github.com/validator-labs/validator-plugin-vsphere/issues/16)) ([0dbadcc](https://github.com/validator-labs/validator-plugin-vsphere/commit/0dbadccbd2e52226022109be8f7cfa28fa948548))
* **deps:** update golang docker tag to v1.21 ([#2](https://github.com/validator-labs/validator-plugin-vsphere/issues/2)) ([212821b](https://github.com/validator-labs/validator-plugin-vsphere/commit/212821bc6443d68d252d4b60e93d4bbeb48b16d0))
* **deps:** update golang:1.21 docker digest to 19600fd ([#19](https://github.com/validator-labs/validator-plugin-vsphere/issues/19)) ([fde96bb](https://github.com/validator-labs/validator-plugin-vsphere/commit/fde96bbe6b54eea6e39d67392c1bfdcb6eb7bf57))
* Minor enhancements to tags and controller ([#32](https://github.com/validator-labs/validator-plugin-vsphere/issues/32)) ([6a79097](https://github.com/validator-labs/validator-plugin-vsphere/commit/6a79097d7e7102b1601d5181aad3e6bbd15c502a))
* refactor and move out functions from vsphere into respective v… ([#12](https://github.com/validator-labs/validator-plugin-vsphere/issues/12)) ([aee8f3d](https://github.com/validator-labs/validator-plugin-vsphere/commit/aee8f3d14bcf53f5a3e818d2011e41a3a05acd5d))
* release 0.0.1 ([#30](https://github.com/validator-labs/validator-plugin-vsphere/issues/30)) ([1ec57dc](https://github.com/validator-labs/validator-plugin-vsphere/commit/1ec57dc549e22b4f6f1bd71eb2e1ea0b9e196588))
* Remove unused RuleEngine struct ([#13](https://github.com/validator-labs/validator-plugin-vsphere/issues/13)) ([c3a1f95](https://github.com/validator-labs/validator-plugin-vsphere/commit/c3a1f95e111ed67e22b45d5ee164dda15734c533))
