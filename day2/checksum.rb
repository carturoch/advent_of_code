class Checksum
  class << self
    def line_diff(line)
      return 0 if line.empty?
      line.max - line.min
    end
  end
end