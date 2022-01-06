## wice completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	wice completion zsh > "${fpath[1]}/_wice"

#### macOS:

	wice completion zsh > /usr/local/share/zsh/site-functions/_wice

You will need to start a new shell for this setup to take effect.


```
wice completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [wice completion](wice_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 6-Jan-2022