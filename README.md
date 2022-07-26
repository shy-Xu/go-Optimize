# go-optimize
> optimize some offical package of golang

## Usage of testflag
>  -test.bench regexp<br>
        run only benchmarks matching regexp<br>
  -test.benchmem<br>
        print memory allocations for benchmarks<br><br>
  -test.benchtime d<br>
        run each benchmark for duration d (default 1s)<br>
  -test.blockprofile file<br>
        write a goroutine blocking profile to file<br>
  -test.blockprofilerate rate<br>
        set blocking profile rate (see runtime.SetBlockProfileRate) (default 1)<br>
  -test.count n<br>
        run tests and benchmarks n times (default 1)<br>
  -test.coverprofile file<br>
        write a coverage profile to file<br>
  -test.cpu list<br>
        comma-separated list of cpu counts to run each test with<br>
  -test.cpuprofile file<br>
        write a cpu profile to file<br>
  -test.failfast<br>
        do not start new tests after the first test failure<br>
  -test.fuzz regexp<br>
        run the fuzz test matching regexp<br>
  -test.fuzzcachedir string<br>
        directory where interesting fuzzing inputs are stored (for use only by cmd/go)<br>
  -test.fuzzminimizetime value<br>
        time to spend minimizing a value after finding a failing input (default 1m0s)<br>
  -test.fuzztime value<br>
        time to spend fuzzing; default is to run indefinitely<br>
  -test.fuzzworker<br>
        coordinate with the parent process to fuzz random values (for use only by cmd/go)<br>
  -test.list regexp<br>
        list tests, examples, and benchmarks matching regexp then exit<br>
  -test.memprofile file<br>
        write an allocation profile to file<br>
  -test.memprofilerate rate<br>
        set memory allocation profiling rate (see runtime.MemProfileRate)<br>
        write a mutex contention profile to the named file after execution<br>
  -test.mutexprofilefraction int<br>
        if >= 0, calls runtime.SetMutexProfileFraction() (default 1)<br>
  -test.outputdir dir<br>
        write profiles to dir<br>
  -test.paniconexit0<br>
        panic on call to os.Exit(0)<br>
  -test.parallel n<br>
  -test.run regexp<br>
        run only tests and examples matching regexp<br>
        run smaller test suite to save time<br>
  -test.shuffle string<br>
        randomize the execution order of tests and benchmarks (default "off")<br>
  -test.testlogfile file<br>
        write test action log to file (for use only by cmd/go)<br>
  -test.timeout d<br>
        panic test binary after duration d (default 0, timeout disabled)<br>
  -test.trace file<br>
        write an execution trace to file<br>
  -test.v<br>
        verbose: print additional output<br>
