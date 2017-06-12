require 'awspec'

name = attribute 'name', {}

control 'default' do
  describe "the instance #{name}" do
    subject { ec2(name) }

    it { should exist }
  end
end
