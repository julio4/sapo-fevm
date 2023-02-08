type JobSummary = {
  id: number;
  address: `0x${string}`;
  status: number | null; // 0: pending, 1: success, 2: fail
  jobId: string | null; // JobId
  exitCode: number;
  outputs: string[]; // CIDs
  stderr: string;
  stdout: string;
  outputUrl: string | null;
};

export default JobSummary;