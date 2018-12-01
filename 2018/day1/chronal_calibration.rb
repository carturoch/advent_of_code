class ChronalCalibration
  attr_accessor :result
  attr_accessor :detected_frequencies

  def initialize
    self.result = 0
    self.detected_frequencies = [0]
    self.duplicated = nil
  end

  def apply(str, stop_on_repeated: false)
    return self.duplicated if duplicated?
    
    self.result += str.to_i
    check_duplication
  end

  def duplicated?
    !self.duplicated.nil?
  end

protected

  attr_accessor :duplicated

  def check_duplication
    if detected_frequencies.include? self.result
      self.duplicated = self.result
    else
      self.detected_frequencies << self.result
    end
  end

end

# ## STEP 1
# 
# Simpler solution
# 
# p File
#   .read('input.txt')
#   .split("\n")
#   .reduce(0) { |memo, v| memo + v.to_i }
# ----------------------------------------------------
# OO way
# 
# device = ChronalCalibration.new
# IO.foreach('input.txt') { |v| device.apply(v) }
# p device.result 
