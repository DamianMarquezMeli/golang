# RUN

go run merge-slices.go

## OUTPUT

Result:
Anna 52
Jane 33
John 31
Mary 41
Paul 24
Tom 66
********************************************************
Expected result:
Anna 52
Jane 33
John 31
Mary 41
Paul 24
Tom 66

# RUN TESTS

go test -v

## OUTPUT

=== RUN   TestMergePointersSlices
    merge-slices_test.go:37: Merge pointers slices: Slices are equal
--- PASS: TestMergePointersSlices (0.00s)
PASS
ok      github.com/devpablocristo/merge-slices  0.002s