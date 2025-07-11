// swift-tools-version: 5.7

import PackageDescription

let package = Package(
  name: "LibXray",
  platforms: [.iOS(.v15), .macOS(.v13)],
  products: [
    .library(name: "LibXray", targets: ["LibXray"])
  ],
  targets: [
    .binaryTarget(
      name: "LibXray",
      url: "https://github.com/tozik/BuildLibXray/releases/download/25.3.30/LibXray.xcframework.zip",
      checksum: "610e372eea160548d42095662bdfe136e394e9f7f8bbe267b79fab37dd9c4efd"
    )
  ]
)
