require './chronal_calibration'

describe ChronalCalibration do

  subject { described_class.new }

  describe 'initialize' do
    it 'result is zero by default' do
      expect(subject.result).to eq 0
    end

    it 'sets a hash to check for duplicated' do
      expect(subject.detected_frequencies).to eq [0]
    end

    it 'sets duplicated flag to false by default' do
      expect(subject).not_to be_duplicated
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

    context 'when considering duplicated frequencies' do
      it 'sets the flag when detects duplicated' do
        subject.apply('+1', stop_on_repeated: true)
        subject.apply('-1', stop_on_repeated: true)

        expect(subject).to be_duplicated
        expect(subject.result).to eq 0
      end
    end
  end
end

