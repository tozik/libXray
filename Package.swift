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
      checksum: "ed34ea1f3348eb03d1239f2c7d59f2c176c2e0caa8ce0ebc87dd8c7d079072ac"
    )
  ]
)
