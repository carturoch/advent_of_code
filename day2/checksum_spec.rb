require './checksum.rb'

describe 'Checksum' do
  describe 'difference in a line' do
    it { expect(Checksum.line_diff([])).to eq 0 }
    it { expect(Checksum.line_diff([5,5,5])).to eq 0 }
    it { expect(Checksum.line_diff([1,5,5,9])).to eq 8 }
    it { expect(Checksum.line_diff([9,5,5,2])).to eq 7 }
  end

  describe 'checksum calculation for a matrix' do
    let(:m) do
      [
        [5, 1, 9, 5],
        [7, 5, 3],
        [2, 4, 6, 8],
      ]
    end

    it { expect(Checksum.of_matrix([[]])).to eq 0 }
    it { expect(Checksum.of_matrix(m)).to eq 18 }
  end
end