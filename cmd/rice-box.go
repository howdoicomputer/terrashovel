package cmd

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "Gemfile",
		FileModTime: time.Unix(1497267295, 0),
		Content:     string("# frozen_string_literal: true\nsource 'https://rubygems.org'\n\ngroup :testing do\n  gem 'awspec'\n  gem 'kitchen-terraform'\n  gem 'rake'\nend\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "README.md",
		FileModTime: time.Unix(1497238646, 0),
		Content:     string(""),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "gitignore",
		FileModTime: time.Unix(1497238646, 0),
		Content:     string(".planfiles\n.terraform\n.kitchen"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "kitchen.yml",
		FileModTime: time.Unix(1497267239, 0),
		Content:     string("---\ndriver:\n  name: terraform\n\nplatforms:\n  - name: aws\n\nsuites:\n  - name: default\n    verifier:\n      name: terraform\n      format: doc\n      groups:\n        - name: default\n          attributes:\n            name: name\n    provisioner:\n      name: terraform\n      apply_timeout: 600\n      color: true\n      directory: >-\n        test/fixtures/default\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "main.tf",
		FileModTime: time.Unix(1497267163, 0),
		Content:     string("data \"aws_ami\" \"ubuntu\" {\n  most_recent = true\n\n  filter {\n    name   = \"name\"\n    values = [\"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*\"]\n  }\n\n  filter {\n    name   = \"virtualization-type\"\n    values = [\"hvm\"]\n  }\n\n  owners = [\"099720109477\"] # Canonical\n}\n\nresource \"aws_instance\" \"foo\" {\n  instance_type = \"t2.micro\"\n  ami           = \"${data.aws_ami.ubuntu.id}\"\n\n  tags {\n    Name    = \"${var.name}\"\n    Testing = \"${var.testing}\"\n  }\n}\n\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "outputs.tf",
		FileModTime: time.Unix(1497267186, 0),
		Content:     string("output \"aws_instance_name\" {\n  value = \"${aws_instance.foo.tags.Name}\"\n}\n"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    "test/fixtures/default/module.tf",
		FileModTime: time.Unix(1497267111, 0),
		Content:     string("provider \"aws\" {\n  region = \"us-east-1\"\n}\n\nresource \"random_id\" \"server\" {\n  byte_length = 8\n}\n\nmodule \"default\" {\n  name    = \"doc_holiday_${random_id.server.hex}\"\n  source  = \"../../../\"\n  testing = \"true\"\n}\n\n"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    "test/fixtures/default/outputs.tf",
		FileModTime: time.Unix(1497267130, 0),
		Content:     string("output \"name\" {\n  value = \"${module.default.aws_instance_name}\"\n}\n"),
	}
	fileg := &embedded.EmbeddedFile{
		Filename:    "test/integration/default/controls/ec2_server_spec.rb",
		FileModTime: time.Unix(1497238646, 0),
		Content:     string("require 'awspec'\n\nname = attribute 'name', {}\n\ncontrol 'default' do\n  describe \"the instance #{name}\" do\n    subject { ec2(name) }\n\n    it { should exist }\n  end\nend\n"),
	}
	fileh := &embedded.EmbeddedFile{
		Filename:    "test/integration/default/controls/state_file_spec.rb",
		FileModTime: time.Unix(1497238646, 0),
		Content:     string("terraform_state = attribute 'terraform_state', {}\n\ncontrol 'state_file' do\n  describe 'the Terraform state file' do\n    subject { json(terraform_state).terraform_version }\n\n    it('is accessible') { is_expected.to match(/\\d+\\.\\d+\\.\\d+/) }\n  end\nend\n"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    "variables.tf",
		FileModTime: time.Unix(1497267206, 0),
		Content:     string("variable \"name\" {\n  description = \"The name of the EC2 instance.\"\n}\n\nvariable \"testing\" {\n  description = \"An indication of whether or not the created infrastructure is meant for testing.\"\n}\n\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1497267295, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "Gemfile"
			file3, // "README.md"
			file4, // "gitignore"
			file5, // "kitchen.yml"
			file6, // "main.tf"
			file7, // "outputs.tf"
			filei, // "variables.tf"

		},
	}
	dir8 := &embedded.EmbeddedDir{
		Filename:   "test",
		DirModTime: time.Unix(1497238646, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir9 := &embedded.EmbeddedDir{
		Filename:   "test/fixtures",
		DirModTime: time.Unix(1497238646, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dira := &embedded.EmbeddedDir{
		Filename:   "test/fixtures/default",
		DirModTime: time.Unix(1497267130, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			fileb, // "test/fixtures/default/module.tf"
			filec, // "test/fixtures/default/outputs.tf"

		},
	}
	dird := &embedded.EmbeddedDir{
		Filename:   "test/integration",
		DirModTime: time.Unix(1497238646, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dire := &embedded.EmbeddedDir{
		Filename:   "test/integration/default",
		DirModTime: time.Unix(1497238646, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirf := &embedded.EmbeddedDir{
		Filename:   "test/integration/default/controls",
		DirModTime: time.Unix(1497238646, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			fileg, // "test/integration/default/controls/ec2_server_spec.rb"
			fileh, // "test/integration/default/controls/state_file_spec.rb"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir8, // "test"

	}
	dir8.ChildDirs = []*embedded.EmbeddedDir{
		dir9, // "test/fixtures"
		dird, // "test/integration"

	}
	dir9.ChildDirs = []*embedded.EmbeddedDir{
		dira, // "test/fixtures/default"

	}
	dira.ChildDirs = []*embedded.EmbeddedDir{}
	dird.ChildDirs = []*embedded.EmbeddedDir{
		dire, // "test/integration/default"

	}
	dire.ChildDirs = []*embedded.EmbeddedDir{
		dirf, // "test/integration/default/controls"

	}
	dirf.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../templates`, &embedded.EmbeddedBox{
		Name: `../templates`,
		Time: time.Unix(1497267295, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":                                  dir1,
			"test":                              dir8,
			"test/fixtures":                     dir9,
			"test/fixtures/default":             dira,
			"test/integration":                  dird,
			"test/integration/default":          dire,
			"test/integration/default/controls": dirf,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"Gemfile":                                              file2,
			"README.md":                                            file3,
			"gitignore":                                            file4,
			"kitchen.yml":                                          file5,
			"main.tf":                                              file6,
			"outputs.tf":                                           file7,
			"test/fixtures/default/module.tf":                      fileb,
			"test/fixtures/default/outputs.tf":                     filec,
			"test/integration/default/controls/ec2_server_spec.rb": fileg,
			"test/integration/default/controls/state_file_spec.rb": fileh,
			"variables.tf":                                         filei,
		},
	})
}
