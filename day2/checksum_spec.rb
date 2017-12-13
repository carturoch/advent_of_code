require './checksum.rb'

describe 'Checksum' do
  describe 'difference in a line' do
    it { expect(Checksum.line_diff([])).to eq 0 }
    it { expect(Checksum.line_diff([5,5,5])).to eq 0 }
    it { expect(Checksum.line_diff([1,5,5,9])).to eq 8 }
    it { expect(Checksum.line_diff([9,5,5,2])).to eq 7 }
  end
end