import subprocess
import os

print("The input, if given with --inputs option is in /inputs")
print("In this case, input folder contains this:")
print()

cmd = "tree /inputs"
process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
out, err = process.communicate()

if err == None:
    print("=================")
    print(out.decode("utf8"))
    print("=================")
else:
    print("=================")
    print(err.decode("utf8"))
    print("=================")

print()
print("Similarly, output goes to /outputs for example, let's add a fun file")

outputDir = '/outputs/'

if not os.path.exists(outputDir):
    os.makedirs(outputDir)

outputFileName = "funfile.txt"
outputFile = os.path.join(outputDir, outputFileName)
with open(outputFile, "w") as f:
    f.write("this was a fun file to read")

print("See /output cid by running the following command:")
print()
print("bacalhau list $JOB_ID --output=json | jq '.[0].Status.JobState.Nodes[] | .Shards.\"0\" | select(.RunOutput)'")