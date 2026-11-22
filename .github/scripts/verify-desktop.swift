import CryptoKit
import Foundation

func readBase64(_ path: String) throws -> Data {
    let text = try String(contentsOfFile: path, encoding: .utf8)
    guard let bytes = Data(base64Encoded: text.trimmingCharacters(in: .whitespacesAndNewlines)) else {
        throw CocoaError(.fileReadCorruptFile)
    }
    return bytes
}
let args = CommandLine.arguments
let key = try Curve25519.Signing.PublicKey(rawRepresentation: readBase64(args[1]))
let archive = try Data(contentsOf: URL(fileURLWithPath: args[3]), options: .mappedIfSafe)
guard try key.isValidSignature(readBase64(args[2]), for: archive) else {
    throw CocoaError(.fileReadCorruptFile)
}
