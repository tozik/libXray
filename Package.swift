// swift-tools-version: 5.7
import PackageDescription

let package = Package(
  name: "XrayKit",
  platforms: [.iOS(.v15), .macOS(.v12)],
  products: [
    .library(name: "LibXrayCustom", targets: ["LibXray"])
  ],
  targets: [
    .binaryTarget(
      name: "LibXray",
      url: "https://github.com/tozik/libXray/releases/download/1.8.9/LibXray.xcframework.zip"",
      checksum: "bbf4c31dd8d54a9888f8844c26262e46224448e804191137a4a7b06ab4be68b6"
    )
  ]
)
