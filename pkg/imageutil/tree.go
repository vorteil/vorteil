package imageutil

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vorteil/vorteil/pkg/vdecompiler"
)

//TreeCMD list contents of directories in a tree-like format.
var TreeCMD = &cobra.Command{
	Use:   "tree IMAGE [FILEPATH]",
	Short: "List contents of directories in a tree-like format.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		flagOs, err := cmd.Flags().GetBool("vpartition")
		if err != nil {
			panic(err)
		}
		img := args[0]

		iio, err := vdecompiler.Open(img)
		if err != nil {
			log.Errorf("%v", err)
			os.Exit(1)
		}
		defer iio.Close()

		var fpath string
		if len(args) > 1 {
			fpath = args[1]
		} else {
			fpath = "/"
		}

		if flagOs {
			if fpath != "" && fpath != "/" && fpath != "." {
				log.Errorf("bad FILE_PATH for vpartition: %s", fpath)
				return
			}

			kfiles, err := iio.KernelFiles()
			if err != nil {
				log.Errorf("%v", err)
				os.Exit(1)
			}

			log.Printf(fpath)

			for i := 0; i < len(kfiles)-1; i++ {
				log.Printf("├── %s", kfiles[i].Name)
			}

			log.Printf("└── %s", kfiles[len(kfiles)-1].Name)
			return
		}

		var code []byte

		var recurse func(int, string) error
		recurse = func(ino int, name string) error {
			inode, err := iio.ResolveInode(ino)
			if err != nil {
				return err
			}

			prefix := ""
			idx := len(code) - 1

			for i, c := range code {
				switch c {
				case 0:
					prefix += "    "
				case 1:
					prefix += "│   "
				case 2:
					if i == idx {
						prefix += "├── "
					} else {
						prefix += "│   "
					}
				case 3:
					if i == idx {
						prefix += "└── "
					} else {
						prefix += "    "
					}
				}
			}

			log.Printf("%s%s", prefix, name)

			if !inode.IsDirectory() {
				return nil
			}

			entries, err := iio.Readdir(inode)
			if err != nil {
				return err
			}

			if len(entries) > 2 {
				idx++
				code = append(code, 2)

				for i := 2; i < len(entries)-1; i++ {
					err = recurse(entries[i].Inode, entries[i].Name)
					if err != nil {
						return err
					}
				}

				code[idx] = 3
				err = recurse(entries[len(entries)-1].Inode, entries[len(entries)-1].Name)
				if err != nil {
					return err
				}

				code = code[:idx]
				idx--
			}

			return nil
		}

		ino, err := iio.ResolvePathToInodeNo(fpath)
		if err != nil {
			log.Errorf("%v", err)
			os.Exit(1)
		}

		err = recurse(ino, fpath)
		if err != nil {
			log.Errorf("%v", err)
			os.Exit(1)
		}
	},
}