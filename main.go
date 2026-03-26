package main    

import (
"os"
"fmt"
"path/filepath"

"time"
"math/rand/v2"
"github.com/fatih/color"
"github.com/briandowns/spinner"
"github.com/schollz/progressbar/v3"
)

const version = "0.1.1"

func main() {
	err := run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func run() error { 

    if len(os.Args) > 1 {
        switch os.Args[1] {
        case "-v", "--version":
            Version()
        case "-h", "--help":
            Usage()
        default:
            Usage()
			return fmt.Errorf("unknown argument: %s", os.Args[1])
        } 
    } else {
        Run()
    }

	return nil
}

func Version() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
}

func Usage() {
    fmt.Println("Usage:", filepath.Base(os.Args[0]), "[options]")
    fmt.Println("Options:")
    fmt.Println("  -v, --version    Print version information and exit")
    fmt.Println("  -h, --help       Print this message and exit")
}

func Run() {

	fmt.Println()
	fmt.Println("Welcome to", color.CyanString("GNDN"), "the work simulator!")
	fmt.Println("version", version)
	fmt.Println("Simulating work with progress bars and random messages...")
	fmt.Println()

	color.Cyan("Press Ctrl+C to stop")
	fmt.Println("--------------------")
	fmt.Println()

	generate_spinner("Starting up")
	initialize()
	update()

	// FIRST


	for {
		do_work(" " + get_random_verb() + " " + get_random_noun() + ":")
		generate_random_messages()
		time.Sleep(2 * time.Second)
	}
}


func initialize() {

	fmt.Println()
	fmt.Println("Initializing...")
	time.Sleep(1 * time.Second)
	fmt.Println()

	stages := []string{
		"Loading configuration",
		"Checking system resources",
		"Checking dependencies",
		"Connecting to network",
		"Connecting to database",
		"Checking security",
		"Connecting to cache",
		"Connecting to message queue",
		"Connecting to storage",
		"Checking integrity",
		"Checking consistency",
	}

	for _, s := range(stages) {

		txt := s+"..."
		msg := fmt.Sprintf("%-30s", txt)
		success := fmt.Sprintf("%-30s %10s", msg, color.GreenString("ok"))
		spin(msg, success, 11, "cyan", 1)
	}

	fmt.Println()
	fmt.Println("Initialization complete")
	fmt.Println()
}

func update() {

	fmt.Println()
	spin("Checking for updates...", "Checking for updates..."+color.CyanString("Done"), 11, "red", 3)
	fmt.Println()

	num := rand.IntN(25)
	fmt.Println("Found ", num, "updates")
	time.Sleep(1 * time.Second)
	fmt.Println()

	for i := range(num) {
		n := i + 1
		txt := fmt.Sprintf("Downloading update %d", n)
		msg := fmt.Sprintf("%-25s", txt)
		suc := ""

		dspin(
			color.CyanString(msg),
			suc,
			35,
			"cyan",
			1,
		)

		txt = fmt.Sprintf("Installing update %d", n)
		msg = fmt.Sprintf("%-25s", txt)
		txt = fmt.Sprintf("Update %d installed\n", n)
		suc = fmt.Sprintf("%-25s", txt)

		dspin(
			color.MagentaString(msg),
			suc,
			35,
			"magenta",
			1,
		)	
	}

	fmt.Println()
	spin("Finalizing...", "Finalizing..."+color.CyanString("Done"), 11, "red", 3)
	fmt.Println()
	
}

func do_work(message string) {

	fmt.Println()
	fmt.Println(message)

	bar := progressbar.Default(100)
	
	for range(100) {

		step := 1
		sleep := 80 + time.Duration(rand.IntN(500)) * time.Millisecond

		bar.Add(step)
		time.Sleep(sleep)
	}
	fmt.Println()
}

func get_random_verb() string {

	verbs := []string{
		"Processing", 
		"Calculating", 
		"Loading", 
		"Analyzing",
		"Updating",
		"Reticulating",
		"Simulating",
		"Optimizing",
		"Compiling",
		"Executing",
		"Rendering",
		"Transforming",
		"Indexing",
		"Synchronizing",
		"Compressing",
		"Encrypting",
		"Decrypting",
		"Transcoding",
		"Parsing",
		"Generating",
		"Evaluating",
		"Running",
	}
	return verbs[rand.IntN(len(verbs))]
}

func get_random_noun() string {
	
	nouns := []string{
		"data",
		"files",
		"images",
		"logs",
		"documents",
		"records",
		"entries",
		"objects",
		"classes",
		"functions",
		"modules",
		"components",
		"services",
		"resources",
		"libraries",
		"packages",
		"items",
		"tasks",
		"jobs",
		"processes",
		"operations",
		"splines",
		"algorithms",
		"models",
		"threads",
		"requests",
		"responses",
	}
	return nouns[rand.IntN(len(nouns))]
}

func get_random_status() string {

	statuses := []string{
		color.GreenString("Done"),
		color.GreenString("Ok"),
		color.GreenString("Good"),
		color.GreenString("Finished"),
		color.GreenString("Completed"),
		color.GreenString("Verified"),
		color.YellowString("Halted"),
		color.YellowString("Stopped"),
		color.RedString("Failed"),
		color.RedString("Error"),
		color.CyanString("Waiting"),
		color.CyanString("Pending"),
		color.CyanString("Updating"),
		color.CyanString("Configuring"),
		color.MagentaString("Unknown"),
		color.MagentaString("Unavailable"),
		color.MagentaString("N/A"),
		color.BlueString("Running"),
		color.BlueString("Processing"),
	}
	return statuses[rand.IntN(len(statuses))]
}

func get_random_check() string {

	checks := []string{
		"Termination Check:",
		"Validation Check:",
		"Integrity Check:",
		"Consistency Check:",
		"Dependency Check:",
		"Cache Validation Check",
		"System Check:",
		"Resource Check:",
		"Security Check:",
		"Performance Check:",
		"Health Check:",
	}
	return checks[rand.IntN(len(checks))]
}


func get_random_status_check() string {

	return get_random_check() + " " + get_random_status()
	
}

func generate_status_check() {
	
	check := get_random_check()
	result := get_random_status()
	fmt.Println()
	quick_spin(check, result)
	fmt.Println()
}

func get_random_statement() string {
	
	statements := []string{
		"Advancing to next stage...",
		"Proceeding to the next step...",
		"Continuing to the next phase...",
		"Progressing to the next level...",
		"Moving on to the next job...",
		"Moving forward...",
		"Continuing...",
		"Advancing to next working set...",
		"Moving to next stage...",
		"Moving to next step...",
		"Moving to next phase...",
		"Moving to next level...",
		"Advancing to next step...",
		"Next phase...",
		"Advancing to Next cluster...",
		"Backing up to previous stage...",
		"Rolling back to previous stage...",
		"Rolling forward to next stage...",
		"Reconfiguring for next stage...",
		"Staging...",
		"Staging next batch...",
		"Staging next group...",
		"Staging next set...",
		"Substaging...",
		"Unstaging...", 
		"Restaging...",
		"Restoring...",
		"Reverting...",
		"Updating...",
		"Setting up...",
		"Cleaning up...",
		"Finalizing...",
		"Optimizing...",
		"Optimizing configuration...",
		"Reloading configuration...",
		"Verifying configuration...",
		"Stabilizing...",
		"Stabilizing the configuration...",
		"Reinitializing...",
		"Rolling back...",
		"Rolling forward...",
		"Reconfiguring...",
		"Resetting...",
		"Tearing down and restarting...",
		"Doing cleanup...",
		"Testing...",
		"Finishing up...",
		"Please wait...",
		"Tidying up...",
		"Stopping...",
		"Restarting...",
		"Flushing cache...",
		"Rebuilding...",
		"Finalizing results...",
		"Cleaning up resources...",
		"Saving progress...",
		"Checking for updates...",
		"Verifying integrity...",
		"Optimizing performance...",
		"Completing tasks...",
	}
	return statements[rand.IntN(len(statements))]
}


func generate_random_statement() {
	msg := get_random_statement()
	fmt.Println()
	spin(
		msg,
		msg+" Done",
		9,
		"white",
		1,
	)
}

func get_random_error() string {
	
	errors := []string{
		"Unable to connect to database",
		"File not found",
		"Access denied",
		"Out of memory",
		"Invalid configuration",
		"Null pointer exception",
		"Segmentation fault",
		"Index out of range",
		"Stack overflow",
		"Invalid input",
		"Timeout occurred",
		"Network unreachable",
		"Permission denied",
		"Disk full",
	}

	code := get_random_number()
	message := color.RedString(
		fmt.Sprintf("Error %s: %s", code, errors[rand.IntN(len(errors))]))

	return message
}

func get_random_warning() string {
	
	warnings := []string{
		color.YellowString("Warning: Low disk space"),
		color.YellowString("Warning: High memory usage"),
		color.YellowString("Warning: Network latency detected"),
		color.YellowString("Warning: Deprecated API usage"),
		color.YellowString("Warning: Unoptimized code detected"),
		color.YellowString("Warning: Potential security vulnerability"),
		color.YellowString("Warning: Configuration mismatch"),
		color.YellowString("Warning: Resource contention detected"),
	}
	return warnings[rand.IntN(len(warnings))]
}

func get_random_info() string {
	
	infos := []string{
		color.CyanString("Waiting for data..."),
		color.CyanString("Retrying operation..."),
		color.CyanString("Connection unstable..."),
		color.CyanString("Processing halted..."),
		color.CyanString("Configuration reloaded!"),
		color.CyanString("Stabilizing..."),
		color.CyanString("Stabilized."),
		color.CyanString("Reinitializing..."),
		color.CyanString("Reinitialization complete."),
		color.CyanString("Maintenace mode enabled..."),
		color.CyanString("Service restarting..."),
	}
	return infos[rand.IntN(len(infos))]
}

func generate_random_info() {
	msg := get_random_info()
	fmt.Println()
	fmt.Println(msg)
	time.Sleep(1 * time.Second)
	spin(
		color.CyanString("Modifying the environment..."),
		color.CyanString("Changes written to " + get_random_string() + ".env"),
		11,
		"cyan",
		3,
	)
	time.Sleep(1 * time.Second)
	spin(
		color.CyanString("Resetting..."),
		color.CyanString("Reset successful!"),
		11,
		"cyan",
		4,
	)
	fmt.Println()

	initialize()
}


func get_random_string() string {

	charset := "abcdefghijklmnopqrstuvwxyz0123456789"

	length := 3 + rand.IntN(5)

	str := ""
	// randomize characters and append to str

	for range(length) {
		ch := string(charset[rand.IntN(len(charset))])
		str += ch

	}

	return str
}

func get_random_filename() string {

	extensions := []string{
		".dll",
		".log",
		".csv",
		".json",
		".xml",
		".yaml",
		".toml",
		".env",
		".conf",
		".tar",
		".tar.gz",
		".zip",
		".rar",
		".ini",
		".lib",
		".db",
		".a",
		".so",
		".jar",
		".dat",
		".sh",
		".o",
		".exe",
	}

	return get_random_string() + extensions[rand.IntN(len(extensions))]

}


func get_random_file_operation() string {
	
	operations := []string{
		"Reading",
		"Writing",
		"Deleting",
		"Moving",
		"Copying",
		"Creating",
		"Compressing",
		"Encrypting",
		"Decrypting",
		"Inflating",
		"Deflating",
		"Archiving",
	}

	return operations[rand.IntN(len(operations))] + " " + get_random_filename()
}


func generate_file_operation_spam() {

	batch := get_random_number()

	fmt.Println()
	fmt.Println("Initiating file operations for batch", batch)
	fmt.Println()

	number := 3 + rand.IntN(70)
	sleep := 50 + time.Duration(rand.IntN(500)) * time.Millisecond
	
	for range(number) {
		fmt.Println("   ", get_random_file_operation())
		time.Sleep(sleep)
	}

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func get_random_symbol() string {

	chance := rand.IntN(100)
	var sym string 

	symbols := []string{
		color.GreenString("+"),
		color.RedString("x"),
		color.YellowString("!"),
		color.CyanString("*"),
		color.MagentaString("#"),
		color.BlueString("@"),
	}

	if chance < 90 {
		sym = " "
	} else {
		sym = symbols[rand.IntN(len(symbols))]
	}

	return sym
}


func generate_grid() {

	stage := get_random_number()

	fmt.Println()
	fmt.Println("Generating grid for stage", stage)
	fmt.Println()

	spin("Generating...", "", 9, "white", 7)

	// generate 10x10 grid of random symbols with borders

	for i := range(10) {
		for j := range(20) {

			if i == 0 && j == 0 { 
				fmt.Print("+")
			} else if i == 0 && j == 19 {
				fmt.Print("+")
			} else if i == 9 && j == 0 {
				fmt.Print("+")
			} else if i == 9 && j == 19 {
				fmt.Print("+")
			} else if j == 0 || j == 19 {
				fmt.Print("|")
			} else if i == 0 || i == 9 {
				fmt.Print("-")
			} else {
				fmt.Print(get_random_symbol())
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func get_random_number() string {

	// generate random 3 digit number
	num := 100 + rand.IntN(899)
	return fmt.Sprintf("%d", num)
}

func generate_bar_graph() {
	
	stage := get_random_number()
	fmt.Println()
	fmt.Println("Generating report for staging batch", stage)
	fmt.Println()

	spin("Generating...", "", 9, "white", 7)

	// generate 4-6 bars
	// each bar has random height between 1 and 50
	// each bar has random color from the set Red, Greed, Yellow, Blue, Magenta, Cyan
	// bars are labeled A, B, C, D, E, F in order
	// print bars with height and color and height number at the end of the bar

	num_bars := 4 + rand.IntN(4)

	for i := range(num_bars) {	
		height := 1 + rand.IntN(50)
		color_func := color.New(color.FgRed).SprintFunc()
		switch rand.IntN(6) {
		case 0:
			color_func = color.New(color.FgRed).SprintFunc()
		case 1:
			color_func = color.New(color.FgGreen).SprintFunc()
		case 2:
			color_func = color.New(color.FgYellow).SprintFunc()
		case 3:
			color_func = color.New(color.FgBlue).SprintFunc()
		case 4:
			color_func = color.New(color.FgMagenta).SprintFunc()
		case 5:
			color_func = color.New(color.FgCyan).SprintFunc()
		}

		label := string('A' + i)
		bar := ""
		for range(height) {
			bar += "█"
		}
		fmt.Printf("%s: %s %d\n", label, color_func(bar), height)
	}

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func generate_checklist() {
	
	group := get_random_number()
	num_items := 5 + rand.IntN(5)

	fmt.Println()
	fmt.Println("Generating staging checklist for stages 1 -", num_items, "in group", group)
	fmt.Println()

	done := 0
	halted := 0
	failed := 0


	for i := range(num_items) {
		status := get_quick_status()
		if status == color.GreenString("Done") {
			done++
		} else if status == color.YellowString("Halted") {
			halted++
		} else {
			failed++
		}

		quick_spin(fmt.Sprintf("Stage %d:", i+1), status)
	}

	fmt.Println()
	fmt.Printf("%d done, %d halted, %d failed\n", done, halted, failed)
	fmt.Println()
	fmt.Println("Done")
	fmt.Println()

}

func generate_random_messages() {

	number := rand.IntN(10)
	sleep := 50 + time.Duration(rand.IntN(500)) * time.Millisecond

	for range(number) {

		chance := rand.IntN(100)
		
		if chance < 70 {
			generate_common_message_()
		} else {
			get_random_rare_message_generator()()
		}
		time.Sleep(sleep)
	}

}

func get_random_logfile() string {

	return fmt.Sprintf("/var/log/%s.log", get_random_string())
}

func get_random_server() string {

	chance := rand.IntN(100)

	var name string

	types := []string {
		"dt",
		"pr",
		"er",
		"xt",
		"zs",
	}
	
	if chance < 75 {
		name = fmt.Sprintf("gn%s-0%s.local", types[rand.IntN(len(types))],get_random_number())
	} else {
		name = fmt.Sprintf("10.%d.%d.%d", rand.IntN(9), rand.IntN(99), rand.IntN(99))
	}
	
	return name
}

func get_random_domain() string {
	domains := []string{
		"contoso.com",
		"fabrikam.com",
		"altostrat.com",
		"cymbalgroup.com",
		"gndn.local",
	}
	return domains[rand.IntN(len(domains))]
}


func generate_random_error() {
	err := get_random_error()
	fmt.Println()
	fmt.Println(err)
	fmt.Print("\a")
	time.Sleep(1 * time.Second)
	spin(
		color.RedString("Generating error log..."),
		color.RedString("Log generated in " + get_random_logfile()),
		11,
		"red",
		3,
	)
	spin(
		color.RedString("Sending notification..."),
		color.RedString("Notification sent to admin@"+get_random_domain()),
		9,
		"red",
		4,
	)
	fmt.Println()
}

func generate_random_warning() {
	err := get_random_warning()
	fmt.Println()
	fmt.Println(err)
	fmt.Print("\a")
	time.Sleep(1 * time.Second)
	spin(
		color.YellowString("Generating log..."),
		color.YellowString("Log generated in " + get_random_logfile()),
		11,
		"yellow",
		3,
	)
	fmt.Println()
}

func generate_common_message_() {

	chance := rand.IntN(100)

	if chance < 40 {
		generate_random_statement()
	} else if chance >= 40 && chance < 60 {
		generate_status_check()
	} else if chance >= 60 && chance < 70 {
		generate_random_error()
	} else if chance >= 70 && chance < 80 {
		generate_random_warning()
	} else if chance >= 80 && chance < 90 {
		generate_random_info()
	} else {
		generate_spinner("")
	}
}

func get_random_rare_message_generator() func() {
	
	generators := []func(){
		generate_file_operation_spam,
		generate_grid,
		generate_bar_graph,
		generate_checklist,
		generate_report,
		generate_batch_operation_spam,
		generate_network_diagnostics,
		generate_batch_report,
		generate_file_downloader,
		initialize,
		update,
	}
	return generators[rand.IntN(len(generators))]
}

func generate_spinner(message string) {

	var verb string

	if message == "" {
		verb = get_random_verb()
	} else {
		verb = message
	}

	charsets := []int {
		1,
		7,
		9,
		11,
		35,
		41,
	}

	charset := charsets[rand.IntN(len(charsets))]

	s := spinner.New(spinner.CharSets[charset], 100*time.Millisecond) 
	s.Prefix = verb + "... "
	s.FinalMSG = verb + "... " + color.GreenString("Done") + "\n"
	s.Start()

	colors := []string {
		"red",
		"green",
		"yellow",
		"blue",
		"magenta",
		"cyan",
	}

	color := colors[rand.IntN(len(colors))]
	s.Color(color)

	duration := 2 + rand.IntN(5)
	time.Sleep(time.Duration(duration) * time.Second)
	s.Stop()
}

func spin(message string, result string, charset int, color string, duration int) {
	s := spinner.New(spinner.CharSets[charset], 100*time.Millisecond) 
	s.Prefix = message + " "
	s.FinalMSG = result + "\n"
	s.Start()
	s.Color(color)
	time.Sleep(time.Duration(duration) * time.Second)
	s.Stop()
}

func dspin(message string, result string, charset int, color string, duration int) {
	s := spinner.New(spinner.CharSets[charset], 100*time.Millisecond)
	s.Prefix = message
	s.FinalMSG = result
	s.Start()
	s.Color(color)
	time.Sleep(time.Duration(duration) * time.Second)
	s.Stop()
}

func quick_spin(message string, result string) {

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) 
	s.Prefix = message + " "
	s.FinalMSG = message + " " + result + "\n"
	s.Start()

	s.Color("cyan")

	duration := 1 + rand.IntN(3)
	time.Sleep(time.Duration(duration) * time.Second)
	s.Stop()
}

func quick_download() {

	file := get_random_filename()
	message := color.CyanString(fmt.Sprintf("Downloading %s", file))
	result := fmt.Sprintf("Downloaded %s", file)

	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond) 
	s.Prefix = message + " "
	s.FinalMSG = result + "\n"
	s.Start()

	s.Color("cyan")

	duration := 1 + rand.IntN(2)
	time.Sleep(time.Duration(duration) * time.Second)
	s.Stop()
}

func get_quick_status() string {

	chance := rand.IntN(100)

	var status string
	
	if chance < 70 {
		status = color.GreenString("Done")
	} else if chance >= 70 && chance < 85 {
		status = color.YellowString("Halted")
	} else {
		status = color.RedString("Failed")
	}

	return status
}

func get_random_percentage() string {
	percent := rand.IntN(100)
	pctstr := fmt.Sprintf("%d%%", percent)

	if percent < 30 {
		return color.RedString(pctstr)
	} else if percent >= 30 && percent < 50 {
		return color.YellowString(pctstr)
	} else if percent >= 50 && percent < 80 {
		return pctstr
	} else {
		return color.GreenString(pctstr)
	}
}

func generate_report() {

	groups := 1 + rand.IntN(8)
	stage := get_random_number()

	fmt.Println()
	fmt.Println("Generating completion report for groups 1 -", groups, "in stage", stage)
	fmt.Println()

	for group := range(groups) {
		gr := fmt.Sprintf("Group %d", group+1)
		quick_spin(gr + ":", get_random_percentage())
	}
	
	fmt.Println()
	fmt.Println("Done")
	fmt.Println()

}

func get_random_batch_operation() string {
	
	operations := []string{
		"Backing up database",
		"Restoring database",
		"Optimizing database",
		"Indexing database",
		"Reindexing database",
		"Rebuilding database",
		"Setting up database",
		"Cleaning up database",
		"Finalizing database",
		"Rebuilding index",
		"Optimizing index",
		"Deleting index",
		"Reconfiguring index",
		"Setting up index",
		"Cleaning up index",
		"Finalizing index",
		"Flushing cache",
		"Clearing cache",
		"Rebuilding cache",
		"Reinitializing cache",
		"Reconfiguring cache",
		"Resetting cache",
		"Optimizing cache",
		"Reindexing cache",
		"Backing up cache",
		"Restoring cache",
		"Deleting cache",
		"Backing up log file",
		"Restoring log file",
		"Deleting log file",
		"Optimizing log file",
		"Indexing log file",
		"Rebuilding log file",
		"Optimizing log file",
		"Reindexing log file",
		"Setting up log file",
		"Cleaning up log file",
		"Finalizing log file",
		"Backing up dataset",
		"Restoring dataset",
		"Optimizing dataset",
		"Indexing dataset",
		"Rebuilding dataset",
		"Optimizing dataset",
		"Reindexing dataset",
		"Setting up dataset",
		"Cleaning up dataset",
		"Finalizing dataset",
		"Backing up library",
		"Restoring library",
		"Optimizing library",
		"Indexing library",
		"Rebuilding library",
		"Optimizing library",
		"Reindexing library",
		"Setting up library",
		"Cleaning up library",
		"Finalizing library",
		"Setting up environment",
		"Cleaning up environment",
		"Finalizing environment",
		"Optimizing environment",
		"Reconfiguring environment",
		"Resetting environment",
		"Reinitializing environment",
		"Reconfiguring environment",
		"Rolling back environment",
		"Setting up configuration",
		"Cleaning up configuration",
		"Finalizing configuration",
		"Optimizing configuration",
		"Reconfiguring configuration",
		"Resetting configuration",
		"Reinitializing configuration",
		"Reconfiguring configuration",
		"Rolling back configuration",
		"Setting up test",
		"Cleaning up test",
		"Finalizing test",
		"Optimizing test",
		"Reconfiguring test",
		"Resetting test",
		"Running test",
		"Reinitializing test",
		"Reconfiguring test",
		"Rolling back test",
		"Setting up stage",
		"Cleaning up stage",
		"Finalizing stage",
		"Optimizing stage",
		"Reconfiguring stage",
		"Resetting stage",
		"Reinitializing stage",
		"Reconfiguring stage",
		"Rolling back stage",
		"Setting up batch",
		"Cleaning up batch",
		"Finalizing batch",
		"Optimizing batch",
		"Reconfiguring batch",
		"Resetting batch",
		"Reinitializing batch",
		"Reconfiguring batch",
		"Rolling back batch",
		"Setting up job",
		"Cleaning up job",
		"Finalizing job",
		"Optimizing job",
		"Reconfiguring job",
		"Resetting job",
		"Reinitializing job",
		"Reconfiguring job",
		"Rolling back job",
		"Setting up process",
		"Cleaning up process",
		"Finalizing process",
		"Optimizing process",
		"Reconfiguring process",
		"Resetting process",
		"Reinitializing process",
		"Reconfiguring process",
		"Rolling back process",
		"Setting up operation",
		"Cleaning up operation",
		"Finalizing operation",
		"Optimizing operation",
		"Reconfiguring operation",
		"Resetting operation",
		"Reinitializing operation",
		"Reconfiguring operation",
		"Rolling back operation",
		"Setting up task",
		"Cleaning up task",
		"Finalizing task",
		"Optimizing task",
		"Reconfiguring task",
		"Resetting task",
		"Reinitializing task",
		"Reconfiguring task",
		"Rolling back task",
		"Setting up thread",
		"Cleaning up thread",
		"Finalizing thread",
		"Optimizing thread",
		"Reconfiguring thread",
		"Resetting thread",
		"Reinitializing thread",
		"Reconfiguring thread",
		"Rolling back thread",
		"Setting up request",
		"Cleaning up request",
		"Finalizing request",
		"Optimizing request",
		"Reconfiguring request",
		"Resetting request",
		"Reinitializing request",
		"Reconfiguring request",
		"Rolling back request",
		"Setting up response",
		"Cleaning up response",
		"Finalizing response",
		"Optimizing response",
		"Reconfiguring response",
		"Resetting response",
		"Reinitializing response",
		"Reconfiguring response",
		"Rolling back response",
	}

	return operations[rand.IntN(len(operations))]
}

func generate_batch_operation_spam() {

	batch := get_random_number()

	fmt.Println()
	fmt.Println("Initiating operations for batch", batch)
	fmt.Println()

	number := 15 + rand.IntN(150)

	for range(number) {
		sleep := 50 + time.Duration(rand.IntN(800)) * time.Millisecond
		op := get_random_number()
		fmt.Println("   ", get_random_batch_operation(), op)
		time.Sleep(sleep)
	}

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func generate_network_diagnostics() {

	nic := rand.IntN(9)
	fmt.Println()
	fmt.Printf("Running network diagnostics for eth%d...\n", nic)
	fmt.Println()

	spin("Running network diagnostic...", "", 9, "white", 7)
	spin("Generating report...", "", 9, "white", 3)
	
	fmt.Printf("   %15s%12s\n", "Download:", fmt.Sprintf("%dMbps", 100+rand.IntN(899)))
	fmt.Printf("   %15s%12s\n", "Upload:", fmt.Sprintf("%dMbps", 10+rand.IntN(90)))
	fmt.Printf("   %15s%12s\n", "Latency:", fmt.Sprintf("%dms",1+rand.IntN(100)))
	fmt.Printf("   %15s%12s\n", "Jitter:", fmt.Sprintf("%dms", rand.IntN(50)))
	fmt.Printf("   %15s%12s\n", "Ping:", fmt.Sprintf("%dms",1+rand.IntN(100)))
	fmt.Printf("   %15s%12s\n", "Packet Loss:", fmt.Sprintf("%d%%",rand.IntN(10)))

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()

}


func generate_batch_report() {
	
	batch := get_random_number()
	fmt.Println()
	fmt.Println("Generating batch report for cluster", batch)
	fmt.Println()

	spin("Generating...", "", 9, "white", 7)

	processed := 100 + rand.IntN(899)
	indexed := processed/(1+rand.IntN(2))
	encrypted := indexed/2
	uploaded := encrypted/2


	fmt.Printf("   %19s%13s\n", "Data Processed:", fmt.Sprintf("%dMb", processed))
	fmt.Printf("   %19s%13s\n", "Data Indexed:", fmt.Sprintf("%dMb", indexed))
	fmt.Printf("   %19s%13s\n", "Data Encrypted:", fmt.Sprintf("%dMb", encrypted))
	fmt.Printf("   %19s%13s\n", "Data Uploaded:", fmt.Sprintf("%dMb", uploaded))
	fmt.Println()

	files := rand.IntN(199)
	bad := files/5
	indexing := files-bad/2

	fmt.Printf("   %19s%13s\n", "Files Processed:", fmt.Sprintf("%d", files))
	fmt.Printf("   %19s%13s\n", "Files Indexed:", fmt.Sprintf("%d", indexing))
	fmt.Printf("   %19s%13s\n", "Bad Files:", fmt.Sprintf("%d", bad))
	fmt.Println()
	
	conc := rand.IntN(15)
	halted := conc/3
	running := conc-halted

	fmt.Printf("   %19s%13s\n", "Concurrent Jobs:", fmt.Sprintf("%d", conc))
	fmt.Printf("   %19s%13s\n", "Running Jobs:", fmt.Sprintf("%d", running))
	fmt.Printf("   %19s%13s\n", "Halted Jobs:", fmt.Sprintf("%d", halted))
	fmt.Println()

	tm := rand.IntN(99)
	mn := tm/9
	mx := tm*3

	fmt.Printf("   %19s%13s\n", "Avg. Time per File:", fmt.Sprintf("%dms", tm))
	fmt.Printf("   %19s%13s\n", "Fastest Time:", fmt.Sprintf("%dms", mn))
	fmt.Printf("   %19s%13s\n", "Slowest Time:", fmt.Sprintf("%dms", mx))
	fmt.Println()


	fmt.Printf("   %19s%13s\n", "Data Lost:", fmt.Sprintf("%dKb", rand.IntN(299)))
	fmt.Printf("   %19s%13s\n", "Error Rate:", fmt.Sprintf("%d%%", rand.IntN(10)))



	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func generate_file_downloader() {

	batch := get_random_number()
	fmt.Println()
	fmt.Println("Downloading files for batch", batch)
	fmt.Println()

	files := 10 + rand.IntN(20)

	for range(files) {
		quick_download()
	}

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func generate_server_report() {
	
	cluster := get_random_number()
	fmt.Println()
	fmt.Printf("Generating server availability report for cluster %s...\n", cluster)
	fmt.Println()

	number := 5 + rand.IntN(10)

	online := 0
	degraded := 0
	offline := 0

	for range(number) {
		server := get_random_server()

		chance := rand.IntN(100)
		
		if chance < 70 {
			quick_spin(
				fmt.Sprintf("Checking %17s:", server),
				color.GreenString("Online"),
			)
			online++
		} else if chance >= 70 && chance < 85 {
			quick_spin(
				fmt.Sprintf("Checking %17s:", server),
				color.YellowString("Degraded"),
			)
			degraded++
		} else {
			quick_spin(
				fmt.Sprintf("Checking %17s:", server),
				color.RedString("Offline"),
			)
			offline++
		}
	}

	fmt.Println()
	fmt.Printf("%d online, %d degraded, %d offline\n", online, degraded, offline)
	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}
