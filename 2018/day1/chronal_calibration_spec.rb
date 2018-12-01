require './chronal_calibration'

describe ChronalCalibration do

  subject { described_class.new }

  describe 'initialize' do
    it 'result is zero by default' do
      expect(subject.result).to eq 0
    end
  end

  describe 'apply' do
    it 'applies positive numbers' do
      subject.apply('+1')
      expect(subject.result).to eq 1      
    end

    it 'applies negative numbers' do
      subject.apply('-1')
      expect(subject.result).to eq -1      
    end

    it 'keeps state on result' do
      subject.apply('+5')
      subject.apply('+1')
      subject.apply('-4')

      expect(subject.result).to eq 2
    end
  end
end

