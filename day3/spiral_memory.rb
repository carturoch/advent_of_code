class SpiralMemory
  class << self
    def is_square_corner?(pos)
      return false if pos < 3
      Math.sqrt(pos) % 1 == 0
    end

    def nearest_square
  end
end