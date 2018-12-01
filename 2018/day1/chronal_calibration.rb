class ChronalCalibration
  attr_accessor :result

  def initialize
    self.result = 0
  end

  def apply(str)
    self.result += str.to_i
  end
end

# Simpler solution
# 
# p File
#   .read('input.txt')
#   .split("\n")
#   .reduce(0) { |memo, v| memo + v.to_i }

# OO way
# 
# calibrator = ChronalCalibration.new
# IO.foreach('input.txt') { |v| calibrator.apply(v) }
# p calibrator.result 