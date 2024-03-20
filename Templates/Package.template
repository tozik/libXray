// swift-tools-version: 5.7
import PackageDescription

let package = Package(
  name: "libXray",
  platforms: [.iOS(.v15), .macOS(.v12)],
  products: [
    .library(name: "libXray", targets: ["libXray"])
  ],
  dependencies: [
      .package(url: "LibXray.xcframework"),
  ],
  targets: [
      .binaryTarget(
          name: "LibXray",
          // iOS
          path: "LibXray.xcframework/ios-arm64/LibXray.framework"
      ),
      .binaryTarget(
          name: "LibXray",
          // iOS Simulator
          path: "LibXray.xcframework/ios-arm64_x86_64-simulator/LibXray.framework"
      ),
      .binaryTarget(
          name: "LibXray",
          // macOS
          path: "LibXray.xcframework/macos-arm64_x86_64/LibXray.framework"
      ),
  ]
)
