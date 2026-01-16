class Currency < Formula
  desc "Currency conversion CLI tool"
  homepage "https://github.com/NoamArie12/Currency-CLI"
  url "https://github.com/NoamArie12/Currency-CLI/releases/download/v1.0.0/cur"
  sha256 "b369ab35e9c9dfe24322d04cb2ccd23724946d115acd11ef586ec38d2f6ca9c6" # <--- Put the hash from the command above here
  version "1.0.0"

  def install
    # This takes the file 'cur' and installs it as 'currency'
    bin.install "cur"
  end

  test do
    # Simple check to see if it runs
    system "#{bin}/currency", "--help"
  end
end
