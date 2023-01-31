import subprocess
import os

print("The input, if given with --inputs option is in /input")
print("In this case, input folder contains this:")

cmd = "tree /input"
process = subprocess.Popen(cmd.split())
process.communicate()

print()
print("Similarly, output goes to /output for example, let's add a fun file")

outputDir = '/output/'

if not os.path.exists(outputDir):
    os.makedirs(outputDir)

outputFileName = "funfile.txt"
outputFile = os.path.join(outputDir, outputFileName)
with open(outputFile, "w") as f:
    f.write("this was a fun file to read")

print("See /output cid by running the following command:")
print()
print("bacalhau list $JOB_ID --output=json | jq '.[0].Status.JobState.Nodes[] | .Shards.\"0\" | select(.RunOutput)'")