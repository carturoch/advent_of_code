require './spiral_memory'

describe 'Spiral Memory' do
  describe 'is_square' do
    it { expect(SpiralMemory.is_squared_corner?(1)).to be false }
    it { expect(SpiralMemory.is_squared_corner?(2)).to be false }
    it { (3..100).each {|n| expect(SpiralMemory.is_squared_corner?((n * n) + 1)).to be true} }
  end

  describe 'nearest_square' do
    it { (1..4).each {|n| expect(SpiralMemory.nearest_square(n)).to eq 4} }
    it { expect(SpiralMemory.nearest_square(5)).to eq 4 }
    it { (8..12).each {|n| expect(SpiralMemory.nearest_square(n)).to eq 9} }
  end

  describe 'offset_from_corner' do
    it 'gets distance from the nearest square' do
      expect(SpiralMemory.offset_from_corner(7)).to eq 0
      expect(SpiralMemory.offset_from_corner(8)).to eq 1
      expect(SpiralMemory.offset_from_corner(9)).to eq 1
      expect(SpiralMemory.offset_from_corner(10)).to eq 0
      expect(SpiralMemory.offset_from_corner(11)).to eq 1
      expect(SpiralMemory.offset_from_corner(12)).to eq 1
      expect(SpiralMemory.offset_from_corner(13)).to eq 0 
      expect(SpiralMemory.offset_from_corner(14)).to eq 1 
    end
  end

  describe 'nearest_corner' do
    it 'returns the given pos is the next one after a square' do
      (2..20).each { |n| expect(SpiralMemory.nearest_corner(n * n + 1)).to eq n*n + 1 }
    end

    it 'detects non square corners' do
      expect(SpiralMemory.nearest_corner(7)).to eq 7 
      expect(SpiralMemory.nearest_corner(13)).to eq 13 
      expect(SpiralMemory.nearest_corner(21)).to eq 21 
      expect(SpiralMemory.nearest_corner(31)).to eq 31 
    end

    it { [7,13,21,31].each { |n| expect(SpiralMemory.nearest_corner(n+1)).to eq n } }
    it { [13,21,31].each { |n| expect(SpiralMemory.nearest_corner(n-1)).to eq n } }
  end

  describe 'distance_to' do
    it { expect(SpiralMemory.distance_to(1)).to eq 0 }
    it { (2..10).each {|n| expect(SpiralMemory.distance_to(n*n)).to eq n-1} }
    it { (2..10).each {|n| expect(SpiralMemory.distance_to(n*n + 1)).to eq n} }
    it 'decreases the distance from the corner to the center' do
      expect(SpiralMemory.distance_to(9)).to eq 2 
      expect(SpiralMemory.distance_to(10)).to eq 3 
      expect(SpiralMemory.distance_to(11)).to eq 2

      expect(SpiralMemory.distance_to(15)).to eq 2
      expect(SpiralMemory.distance_to(16)).to eq 3
      expect(SpiralMemory.distance_to(17)).to eq 4
      expect(SpiralMemory.distance_to(18)).to eq 3
      expect(SpiralMemory.distance_to(19)).to eq 2
    end

    it 'increases from the center to the corner' do
      expect(SpiralMemory.distance_to(12)).to eq 3
      expect(SpiralMemory.distance_to(23)).to eq 2
    end
  end
end