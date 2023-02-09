type JobSummary = {
  id: number;
  address: `0x${string}`;
  status: number | null; // 0: pending, 1: success, 2: fail
  jobId: string | null; // JobId
  exitCode: number | null;
  outputs: string[]; // CIDs
  stderr: string;
  stdout: string;
  outputUrl: string | null;
  cidStderr: string | null;
  cidStdout: string | null;
  cidOutputs: string | null;
};

export default JobSummary;
