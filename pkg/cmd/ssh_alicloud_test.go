package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func Test_buildSshCommand(t *testing.T) {
	attrs := AliyunInstanceAttribute{
		PrivateIP:      "10.11.12.13",
		BastionSSHUser: "bastionUser",
		BastionIP:      "8.9.10.11",
	}

	sshCommand := buildSshCommand("KEY42", &attrs, "user2")
	fmt.Println(sshCommand)

	expected := `ssh -i KEY42 -o "ProxyCommand ssh -i KEY42 -o StrictHostKeyChecking=no -W 10.11.12.13:22 bastionUser@8.9.10.11" user2@10.11.12.13 -o StrictHostKeyChecking=no`

	if expected != sshCommand {
		t.Error("commands didn't match")
	}
}

func Test_buildSshCommandArgs(t *testing.T) {
	attrs := AliyunInstanceAttribute{
		PrivateIP:      "10.11.12.13",
		BastionSSHUser: "bastionUser",
		BastionIP:      "8.9.10.11",
	}

	sshCommand := buildSshCommandArgs("KEY42", &attrs, "user2")
	fmt.Println(sshCommand)

	expected := `-i KEY42 -o ProxyCommand ssh -i KEY42 -o StrictHostKeyChecking=no -W 10.11.12.13:22 bastionUser@8.9.10.11 user2@10.11.12.13 -o StrictHostKeyChecking=no`

	join := strings.Join(sshCommand, " ")
	fmt.Println(expected)
	fmt.Println(join)
	if expected != join {
		t.Error("commands didn't match")
	}
}

func Test_buildAliyunCommand(t *testing.T) {
	attrs := AliyunInstanceAttribute{
		InstanceID: "instanceIDD",
	}
	command := buildAliyunCommand(&attrs)

	expected := "aliyun ecs DescribeInstanceAttribute --InstanceId=instanceIDD"
	if expected != command {
		t.Error("commands didn't match")
	}
}

func Test_buildAliyunCommandArgs(t *testing.T) {
	attrs := AliyunInstanceAttribute{
		InstanceID: "instanceIDD",
	}
	command := buildAliyunCommandArgs(&attrs)
	fmt.Println(command)

	expected := "ecs DescribeInstanceAttribute --InstanceId=instanceIDD"
	if expected != strings.Join(command, " ") {
		t.Error("commands didn't match")
	}
}

func Test_buildBastionCommand(t *testing.T) {
	expected := "aliyun ecs DescribeSecurityGroups --VpcId=VPCIDD"
	command := buildBastionCommand(&AliyunInstanceAttribute{VpcID: "VPCIDD"})
	if command != expected {
		t.Error("commands didn't match")
	}
}
func Test_buildBastionCommandArgs(t *testing.T) {
	expected := "ecs DescribeSecurityGroups --VpcId=VPCIDD"
	command := buildBastionCommandArgs(&AliyunInstanceAttribute{VpcID: "VPCIDD"})
	if expected != strings.Join(command, " ") {
		t.Error("commands didn't match")
	}
}
