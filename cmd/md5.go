package cmd

import (
	"fmt"
	"os"

	"github.com/Fristiner/certutil/utils"
	"github.com/spf13/cobra"
)

var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "A brief description of your command",
	Long:  ` asd `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: md5CmdRun,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	if cmd.Flags().Changed("test") {
	// 		runTestFunction()
	// 		return
	// 	} else if cmd.Flags().Changed("compute") {
	// 		computeFilePath, err := cmd.Flags().GetString("compute")
	// 		if err != nil || computeFilePath == "" {
	// 			fmt.Println("No file provided for computing MD5. Use -c or --compute flag followed by a file path.")
	// 			return
	// 		}
	//
	// 		byteContent, err := utils.ReadFileToBytesWithMultiThread(computeFilePath)
	// 		if err != nil {
	// 			fmt.Printf("Error reading file: %v\n", err)
	// 			return
	// 		}
	// 		md5Value := utils.GetMD5Value(byteContent)
	// 		fmt.Printf("MD5 value of file '%s' is: %s\n", computeFilePath, md5Value)
	// 		return
	// 	} else if cmd.Flags().Changed("encrypt") {
	// 		sourceFilePath, err := cmd.Flags().GetString("encrypt")
	// 		if err != nil || sourceFilePath == "" {
	// 			fmt.Println("No file provided for computing MD5. Use -h or --hex flag followed by a file path.")
	// 			return
	// 		}
	// 		// sourceFilePath 为传入的文件
	// 		// handleHexFlag(cmd, sourceFilePath)
	// 		utils.HandleHexFlag(sourceFilePath)
	// 		return
	// 	} else if cmd.Flags().Changed("string") {
	// 		stringToEncrypt, err := cmd.Flags().GetString("string")
	// 		if err != nil || stringToEncrypt == "" {
	// 			fmt.Println("No string provided for computing MD5. Use -s or --string flag followed by a string.")
	// 			return
	// 		}
	// 		value := utils.GetMD5Value([]byte(stringToEncrypt))
	// 		fmt.Printf("MD5 value of string '%s' is: %s", stringToEncrypt, value)
	// 		return
	// 	} else {
	// 		displayHelpMessage()
	// 		return
	// 	}
	// },
}

//goland:noinspection t

func md5CmdRun(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("test") {
		runTestFunction()
		return
	} else if cmd.Flags().Changed("compute") {
		computeFilePath, err := cmd.Flags().GetString("compute")
		if err != nil || computeFilePath == "" {
			fmt.Println("No file provided for computing MD5. Use -c or --compute flag followed by a file path.")
			return
		}

		byteContent, err := utils.ReadFileToBytesWithMultiThread(computeFilePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		md5Value := utils.GetMD5Value(byteContent)
		fmt.Printf("MD5 value of file '%s' is: %s\n", computeFilePath, md5Value)
		return
	} else if cmd.Flags().Changed("encrypt") {
		sourceFilePath, err := cmd.Flags().GetString("encrypt")
		if err != nil || sourceFilePath == "" {
			fmt.Println("No file provided for computing MD5. Use -h or --hex flag followed by a file path.")
			return
		}
		// sourceFilePath 为传入的文件
		// handleHexFlag(cmd, sourceFilePath)
		utils.HandleHexFlag(sourceFilePath)
		return
	} else if cmd.Flags().Changed("string") {
		stringToEncrypt, err := cmd.Flags().GetString("string")
		if err != nil || stringToEncrypt == "" {
			fmt.Println("No string provided for computing MD5. Use -s or --string flag followed by a string.")
			return
		}
		value := utils.GetMD5Value([]byte(stringToEncrypt))
		fmt.Printf("MD5 value of string '%s' is: %s", stringToEncrypt, value)
		return
	} else if cmd.Flags().Changed("validate") {
		// getString, err := cmd.Flags().GetString("validate")
		// if err != nil {
		// 	fmt.Println("输入错误")
		// 	return
		// }
		// fmt.Println("validate value is " + getString)
		handleValidateFlag(cmd, args)
		fmt.Println(args)
		return
	} else {
		displayHelpMessage()
		return
	}
}

// 14a9f8c6f825091c7ca23da3bce1dfd8
// asdasdads
func displayHelpMessage() {
	fmt.Println("Flags:")
	fmt.Println("-h, --help help for md5")
}

func runTestFunction() {
	str1 := "hello,world"
	str2 := "a"
	str3 := "message digest"
	str4 := "abcdefghijklmnopqrstuvwxyz"
	helloString := []byte(str1)
	helloString2 := []byte(str2)
	helloString3 := []byte(str3)
	helloString4 := []byte(str4)
	value1 := utils.GetMD5Value(helloString2)
	value2 := utils.GetMD5Value(helloString3)
	value3 := utils.GetMD5Value(helloString4)
	value := utils.GetMD5Value(helloString)
	fmt.Println("MD5 value of hello,world is:", value)
	// fmt.Println("Running the test function for MD5 application.")
	fmt.Println("MD5 value of a is:", value1)
	fmt.Println("MD5 value of message digest is:", value2)
	fmt.Println("MD5 value of abcdefghijklmnopqrstuvwxyz is:", value3)
}

func setHelpMd5(cmd *cobra.Command, args []string) {
	fmt.Println("usage:   [-h] --help information")
	fmt.Println("         [-t] --test MD5 application test")
	fmt.Println("         [-s] --string [The string to get the md5 value]")
	fmt.Println("                  compute MD5 of the given string ")
	fmt.Println("         [-c] --compute [file path of the file computed]")
	fmt.Println("                  compute MD5 of the given file")
	fmt.Println("         [-e] --encrypt [file path]     Output the MD5 hash as a .hex file in the current directory")
	fmt.Println("                  compute the MD5 and save it as a .hex file")
	fmt.Println("         [-v] [file path of the file validated or string ] [md5 value ]")
	fmt.Println("                  --validate the integrality of a given file by manual input MD5 value")
	fmt.Println("         [-f] [file path of the file validated] [file path of the .hex file]")
	fmt.Println("                  --validate the integrality of a given file by read MD5 value from .hex file")
}

//goland:noinspection t
func handleValidateFlag(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: md5 -v <filename_or_string> <md5sum>")
		return
	}
	target, err := cmd.Flags().GetString("validate")
	if err != nil {
		fmt.Println("读取参数错误")
		return
	}
	// target := args[0]
	md5Checksum := args[0]

	// 验证MD5值是否有效
	if !utils.IsValidMD5(md5Checksum) {
		fmt.Println("Invalid MD5 checksum format.")
		return
	}

	// 判断目标是文件还是字符串
	if _, err := os.Stat(target); err == nil {
		// 是文件
		// if filepath.IsAbs(target) || strings.HasPrefix(target, "./") || strings.HasPrefix(target, "../") {
		// 确保是绝对路径或相对路径
		// 计算文件md5的值
		// fileMD5Checksum := calculateFileMD5(target)
		fileMD5Checksum, err := utils.GetMD5ValueFile(target)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}
		if fileMD5Checksum != md5Checksum {
			fmt.Printf("File '%s' has a different MD5 checksum: expected %s, got %s\n", target, md5Checksum, fileMD5Checksum)
		} else {
			fmt.Printf("File '%s' has the correct MD5 checksum: %s\n", target, fileMD5Checksum)
		}
		// } else {
		// 	fmt.Println("The provided file path is invalid.")
		// }
	} else if os.IsNotExist(err) {
		// 不是文件，则视为字符串
		calculatedMD5 := utils.GetMD5Value([]byte(target))
		// calculatedMD5 := calculateMD5([]byte(target))
		if calculatedMD5 != md5Checksum {
			fmt.Printf("String has a different MD5 checksum: expected %s, got %s\n", md5Checksum, calculatedMD5)
		} else {
			fmt.Printf("String has the correct MD5 checksum: %s\n", calculatedMD5)
		}
	} else {
		fmt.Printf("Error checking target: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(md5Cmd)
	md5Cmd.SetHelpFunc(setHelpMd5)
	md5Cmd.Flags().BoolP("test", "t", true, "test MD5 application test")
	md5Cmd.Flags().StringP("compute", "c", " ", "compute MD5 of the given file")
	md5Cmd.Flags().StringP("encrypt", "e", " ", "Compute the MD5 and save it as a .hex file")
	md5Cmd.Flags().StringP("string", "s", "", "The string to get the md5 value")
	md5Cmd.Flags().StringP("validate", "v", "", "Validate the integrality of a file or string using provided MD5 sum")
}
