package subcommand

import "os"

func (sc *Subcommand) copyShellFile() error {
	file, err := os.ReadFile(sc.SourceScript)
	if err != nil {
		return err
	}

	targetFile := sc.Hash.String() + ".sh"
	err = os.WriteFile(targetFile, file, 0666)
	if err != nil {
		return err
	}

	return nil
}
