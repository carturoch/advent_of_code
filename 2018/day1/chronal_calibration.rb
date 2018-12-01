class ChronalCalibration
  attr_accessor :result

  def initialize
    self.result = 0
  end

  def apply(str)
    self.result += str.to_i
  end
end