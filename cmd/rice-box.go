package cmd

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "Gemfile",
		FileModTime: time.Unix(1497260074, 0),
		Content:     string("# frozen_string_literal: true\r\nsource 'https://rubygems.org'\r\n\r\ngroup :testing do\r\n  gem 'kitchen-terraform'\r\n  gem 'awspec'\r\n  gem 'rake'\r\nend\r\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "README.md",
		FileModTime: time.Unix(1497259700, 0),
		Content:     string(""),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "gitignore",
		FileModTime: time.Unix(1497260116, 0),
		Content:     string(".planfiles\r\n.terraform\r\n.kitchen"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "kitchen.yml",
		FileModTime: time.Unix(1497263160, 0),
		Content:     string("---\r\ndriver:\r\n  name: terraform\r\n\r\nplatforms:\r\n  - name: aws\r\n\r\nsuites:\r\n  - name: default\r\n    verifier:\r\n      name: terraform\r\n      format: doc\r\n      groups:\r\n        - name: default\r\n          attributes:\r\n            name: name\r\n          controls:\r\n            - ec2_server_spec\r\n            - state_file_spec\r\n    provisioner:\r\n      name: terraform\r\n      apply_timeout: 600\r\n      color: true\r\n      directory: >-\r\n        test/fixtures/default\r\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "main.tf",
		FileModTime: time.Unix(1497260045, 0),
		Content:     string("resource \"aws_instance\", \"${var.name}\" {\r\n  instance_type = \"t2.micro\"\r\n\r\n  tags {\r\n    Name = \"${var.name}\"\r\n  }\r\n}\r\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "outputs.tf",
		FileModTime: time.Unix(1497259965, 0),
		Content:     string("output \"aws_instance.id\" {\r\n  value = \"${aws_instance.foo.id}\"\r\n}\r\n"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    "test/fixtures/default/module.tf",
		FileModTime: time.Unix(1497260197, 0),
		Content:     string("provider \"aws\" {\r\n  region = \"us-east-1\"\r\n}\r\n\r\nmodule \"default\" {\r\n  name = \"doc-holiday\"\r\n}\r\n"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    "test/fixtures/default/outputs.tf",
		FileModTime: time.Unix(1497260233, 0),
		Content:     string("output \"name\" {\r\n  value = \"${module.default.name}\"\r\n}\r\n"),
	}
	fileg := &embedded.EmbeddedFile{
		Filename:    "test/integration/default/controls/ec2_server_spec.rb",
		FileModTime: time.Unix(1497260356, 0),
		Content:     string("require 'awspec'\r\n\r\nname = attribute 'name', {}\r\n\r\ncontrol 'default' do\r\n  describe \"the instance #{name}\" do\r\n    subject { ec2(name) }\r\n\r\n    it { should exist }\r\n  end\r\nend\r\n"),
	}
	fileh := &embedded.EmbeddedFile{
		Filename:    "test/integration/default/controls/state_file_spec.rb",
		FileModTime: time.Unix(1497260385, 0),
		Content:     string("terraform_state = attribute 'terraform_state', {}\r\n\r\ncontrol 'state_file' do\r\n  describe 'the Terraform state file' do\r\n    subject { json(terraform_state).terraform_version }\r\n\r\n    it('is accessible') { is_expected.to match(/\\d+\\.\\d+\\.\\d+/) }\r\n  end\r\nend\r\n"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    "variables.tf",
		FileModTime: time.Unix(1497260006, 0),
		Content:     string("variable \"name\" {\r\n  description = \"The name of the EC2 instance.\"\r\n}\r\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1497263160, 0),
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
		DirModTime: time.Unix(1497274495, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir9 := &embedded.EmbeddedDir{
		Filename:   "test/fixtures",
		DirModTime: time.Unix(1497274451, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dira := &embedded.EmbeddedDir{
		Filename:   "test/fixtures/default",
		DirModTime: time.Unix(1497260233, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			fileb, // "test/fixtures/default/module.tf"
			filec, // "test/fixtures/default/outputs.tf"

		},
	}
	dird := &embedded.EmbeddedDir{
		Filename:   "test/integration",
		DirModTime: time.Unix(1497274495, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dire := &embedded.EmbeddedDir{
		Filename:   "test/integration/default",
		DirModTime: time.Unix(1497274495, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirf := &embedded.EmbeddedDir{
		Filename:   "test/integration/default/controls",
		DirModTime: time.Unix(1497260385, 0),
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
		Time: time.Unix(1497263160, 0),
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
