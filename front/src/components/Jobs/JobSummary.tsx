type JobSummary = {
  id: number;
  address: `0x${string}`;
  status: number | null; // 0: pending, 1: success, 2: fail
  jobId: string | null; // JobId
  exitCode: string | null;
  outputs: string[]; // CIDs
  stderr: string | null;
  stdout: string | null;
  outputUrl: string | null;
  cidStderr: string | null;
  cidStdout: string | null;
  cidOutputs: string | null;
  outputsNamesAndCids: string[] | null;
};

export default JobSummary;
