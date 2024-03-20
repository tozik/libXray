// swift-tools-version: 5.7
import PackageDescription

let package = Package(
  name: "libXray",
  platforms: [.iOS(.v15), .macOS(.v12)],
  products: [
    .library(name: "libXray", targets: ["libXray"])
  ],
  targets: [
    .binaryTarget(
      name: "libXray",
      url: "https://github.com/tozik/libXray/releases/download/1.8.9/LibXray.xcframework.zip"",
      checksum: "37ed6fac9a5e92988c37a4821f28375168584429e37f5304dd8b412f2de06727"
    )
  ]
)
