require './checksum.rb'

describe 'Checksum' do
  describe 'difference in a line' do
    it { expect(Checksum.line_diff([])).to eq 0 }
    it { expect(Checksum.line_diff([5,5,5])).to eq 0 }
    it { expect(Checksum.line_diff([1,5,5,9])).to eq 8 }
    it { expect(Checksum.line_diff([9,5,5,2])).to eq 7 }
  end

  describe 'evenly dividers' do
    it { expect(Checksum.line_ediv([])).to eq 0 }
    it { expect(Checksum.line_ediv([3,5,7,13])).to eq 0 }
    it 'the first exact division' do
      expect(Checksum.line_ediv([5, 9, 2, 8])).to eq 4
    end
  end

  describe 'checksum calculation for a matrix' do
    let(:m) do
      [
        [5, 1, 9, 5],
        [7, 5, 3],
        [2, 4, 6, 8]
      ]
    end

    it { expect(Checksum.of_matrix([[]])).to eq 0 }
    it { expect(Checksum.of_matrix(m)).to eq 18 }
  end

  describe 'checksum second calculation for a matrix' do
    let(:m) do
      [
        [5, 9, 2, 8],
        [9, 4, 7, 3],
        [3, 8, 6, 5]
      ]
    end

    it { expect(Checksum.of_matrix_2([[]])).to eq 0 }
    it { expect(Checksum.of_matrix_2(m)).to eq 9 }
  end
end
