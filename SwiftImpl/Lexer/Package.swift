// swift-tools-version: 6.4
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
  name: "Lexer",
  products: [
    .library(
      name: "Lexer",
      targets: ["Lexer"]
    )
  ],
  dependencies: [
    .package(path: "../Token")
  ],
  targets: [
    .target(
      name: "Lexer",
      dependencies: [
        .product(name: "Token", package: "Token")
      ]
    ),
    .testTarget(
      name: "LexerTests",
      dependencies: ["Lexer"]
    ),
  ],
  swiftLanguageModes: [.v6]
)
