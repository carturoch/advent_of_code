class SpiralMemory
  class << self
    def is_squared_corner?(pos)
      return false if pos < 3
      Math.sqrt(pos - 1) % 1 == 0
    end

    def nearest_corner(pos)
      return pos if is_squared_corner?(pos)
      nsq = nearest_square(pos)
      side = Math.sqrt(nsq).to_i
      ofs = (nsq + 1 - pos).abs
      return pos if ofs == side
      nnsq = if pos > nsq 
                nsq + 1 + side
             else
                nsq + 1 - side
             end

      return nnsq if ofs > side / 2
      nsq + 1
    end

    def nearest_square(pos)
      return 4 if pos < 4
      Math.sqrt(pos).round ** 2
    end

    def offset_from_corner(pos)
      nc = nearest_corner(pos)
      (nc - pos).abs
    end

    def distance_to(pos)
      return 0 if pos == 1
      corner = nearest_corner(pos)
      ofc = offset_from_corner(pos)
      side = Math.sqrt(nearest_square(pos)).to_i
      unless is_squared_corner?(corner)
        side = side - 1 if corner < pos
        side = side + 1 if corner > pos
      end
      side - ofc
    end
  end
end

# puts SpiralMemory.distance_to(361527)