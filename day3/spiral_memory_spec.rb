require './spiral_memory'

describe 'Spiral Memory' do
  describe 'is_square' do
    it { expect(SpiralMemory.is_square_corner?(1)).to be false }
    it { expect(SpiralMemory.is_square_corner?(2)).to be false }
    it { (3..100).each {|n| expect(SpiralMemory.is_square_corner?(n * n)).to be true} }
  end
  describe 'nearest_square' do
    it { expect(SpiralMemory.nearest_square(1)).to eq 1 }
    it { expect(SpiralMemory.nearest_square(1)).to eq 1 }
  end
end