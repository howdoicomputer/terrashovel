package cmd

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "Dockerfile",
		FileModTime: time.Unix(1510944041, 0),
		Content:     string("FROM howdoicomputer/terraform_testing:0.1\n\nRUN gem install bundler\n\nCOPY . /\n\nRUN bundle install\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "Gemfile",
		FileModTime: time.Unix(1510942025, 0),
		Content:     string("# frozen_string_literal: true\nsource 'https://rubygems.org'\n\ngroup :testing do\n  gem 'awspec'\n  gem 'kitchen-terraform', '0.7.0'\n  gem 'inspec', '1.29.0'\n  gem 'rake'\nend\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "Makefile",
		FileModTime: time.Unix(1510944744, 0),
		Content:     string(".PHONY: test\n\ndefault: | build test\n\nbuild:\n\tdocker build . -t foobar\n\ntest:\n\tdocker run --rm -v ~/.aws:/root/.aws foobar bundle exec kitchen test\n\nenv:\n\tdocker run --rm -v ~/.aws:/root/.aws -it foobar /bin/bash\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "README.md",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string(""),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "gitignore",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string(".planfiles\n.terraform\n.kitchen"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "kitchen.yml",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("---\ndriver:\n  name: terraform\n\nplatforms:\n  - name: aws\n\nsuites:\n  - name: default\n    verifier:\n      name: terraform\n      format: doc\n      groups:\n        - name: default\n          attributes:\n            name: name\n    provisioner:\n      name: terraform\n      apply_timeout: 600\n      color: true\n      directory: >-\n        test/fixtures/default\n"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "main.tf",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("data \"aws_ami\" \"ubuntu\" {\n  most_recent = true\n\n  filter {\n    name   = \"name\"\n    values = [\"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*\"]\n  }\n\n  filter {\n    name   = \"virtualization-type\"\n    values = [\"hvm\"]\n  }\n\n  owners = [\"099720109477\"] # Canonical\n}\n\nresource \"aws_instance\" \"foo\" {\n  instance_type = \"t2.micro\"\n  ami           = \"${data.aws_ami.ubuntu.id}\"\n\n  tags {\n    Name    = \"${var.name}\"\n    Testing = \"${var.testing}\"\n  }\n}\n\n"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    "outputs.tf",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("output \"aws_instance_name\" {\n  value = \"${aws_instance.foo.tags.Name}\"\n}\n"),
	}
	filed := &embedded.EmbeddedFile{
		Filename:    "test/fixtures/default/module.tf",
		FileModTime: time.Unix(1510942064, 0),
		Content:     string("provider \"aws\" {\n  region = \"us-east-1\"\n}\n\nresource \"random_id\" \"server\" {\n  byte_length = 8\n}\n\nmodule \"default\" {\n  name    = \"doc_holiday_${random_id.server.hex}\"\n  source  = \"../../../\"\n  testing = \"true\"\n}\n"),
	}
	filee := &embedded.EmbeddedFile{
		Filename:    "test/fixtures/default/outputs.tf",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("output \"name\" {\n  value = \"${module.default.aws_instance_name}\"\n}\n"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    "test/integration/default/controls/ec2_server_spec.rb",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("require 'awspec'\n\nname = attribute 'name', {}\n\ncontrol 'default' do\n  describe \"the instance #{name}\" do\n    subject { ec2(name) }\n\n    it { should exist }\n  end\nend\n"),
	}
	filej := &embedded.EmbeddedFile{
		Filename:    "test/integration/default/controls/state_file_spec.rb",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("terraform_state = attribute 'terraform_state', {}\n\ncontrol 'state_file' do\n  describe 'the Terraform state file' do\n    subject { json(terraform_state).terraform_version }\n\n    it('is accessible') { is_expected.to match(/\\d+\\.\\d+\\.\\d+/) }\n  end\nend\n"),
	}
	filek := &embedded.EmbeddedFile{
		Filename:    "variables.tf",
		FileModTime: time.Unix(1510941275, 0),
		Content:     string("variable \"name\" {\n  description = \"The name of the EC2 instance.\"\n}\n\nvariable \"testing\" {\n  description = \"An indication of whether or not the created infrastructure is meant for testing.\"\n}\n\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1510944744, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "Dockerfile"
			file3, // "Gemfile"
			file4, // "Makefile"
			file5, // "README.md"
			file6, // "gitignore"
			file7, // "kitchen.yml"
			file8, // "main.tf"
			file9, // "outputs.tf"
			filek, // "variables.tf"

		},
	}
	dira := &embedded.EmbeddedDir{
		Filename:   "test",
		DirModTime: time.Unix(1510941275, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirb := &embedded.EmbeddedDir{
		Filename:   "test/fixtures",
		DirModTime: time.Unix(1510941275, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirc := &embedded.EmbeddedDir{
		Filename:   "test/fixtures/default",
		DirModTime: time.Unix(1510942064, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filed, // "test/fixtures/default/module.tf"
			filee, // "test/fixtures/default/outputs.tf"

		},
	}
	dirf := &embedded.EmbeddedDir{
		Filename:   "test/integration",
		DirModTime: time.Unix(1510941275, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirg := &embedded.EmbeddedDir{
		Filename:   "test/integration/default",
		DirModTime: time.Unix(1510941275, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirh := &embedded.EmbeddedDir{
		Filename:   "test/integration/default/controls",
		DirModTime: time.Unix(1510941275, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filei, // "test/integration/default/controls/ec2_server_spec.rb"
			filej, // "test/integration/default/controls/state_file_spec.rb"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dira, // "test"

	}
	dira.ChildDirs = []*embedded.EmbeddedDir{
		dirb, // "test/fixtures"
		dirf, // "test/integration"

	}
	dirb.ChildDirs = []*embedded.EmbeddedDir{
		dirc, // "test/fixtures/default"

	}
	dirc.ChildDirs = []*embedded.EmbeddedDir{}
	dirf.ChildDirs = []*embedded.EmbeddedDir{
		dirg, // "test/integration/default"

	}
	dirg.ChildDirs = []*embedded.EmbeddedDir{
		dirh, // "test/integration/default/controls"

	}
	dirh.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../templates`, &embedded.EmbeddedBox{
		Name: `../templates`,
		Time: time.Unix(1510944744, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":                                  dir1,
			"test":                              dira,
			"test/fixtures":                     dirb,
			"test/fixtures/default":             dirc,
			"test/integration":                  dirf,
			"test/integration/default":          dirg,
			"test/integration/default/controls": dirh,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"Dockerfile":                                           file2,
			"Gemfile":                                              file3,
			"Makefile":                                             file4,
			"README.md":                                            file5,
			"gitignore":                                            file6,
			"kitchen.yml":                                          file7,
			"main.tf":                                              file8,
			"outputs.tf":                                           file9,
			"test/fixtures/default/module.tf":                      filed,
			"test/fixtures/default/outputs.tf":                     filee,
			"test/integration/default/controls/ec2_server_spec.rb": filei,
			"test/integration/default/controls/state_file_spec.rb": filej,
			"variables.tf":                                         filek,
		},
	})
}
