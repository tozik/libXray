// swift-tools-version: 5.7

import PackageDescription

let package = Package(
  name: "libXray",
  platforms: [.iOS(.v16), .macOS(.v12)],
  products: [
    .library(
        name: "libXray",
        targets: ["libXray"]
    )
  ],
  targets: [
    .target(
        name: "libXray",
        dependencies: ["libXray"]
    ),
    .target(
        name: "libXray",
        publicHeadersPath: "."
    ),
    .binaryTarget(
        name: "libXray",
        url: "https://github.com/tozik/libXray/releases/download/1.8.9/LibXray.xcframework.zip",
        checksum: "32a869804e45c52f4c7a3d6a898bb28d289d6b26d8043cb1340bf6f4b78b8f8b"
    )
  ]
)