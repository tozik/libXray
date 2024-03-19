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
        url: "https://github.com/tozik/libXray/releases/download/1.0/LibXray.xcframework.zip",
        checksum: "c402723e6441d0844b97997d8985aa9b52aef6a9263a21e6a7ebff344d8e7405"
    )
  ]
)