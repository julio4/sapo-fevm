type JobResult = {
  id: number;
  address: string;
  status: number; // 0: pending, 1: success, 2: fail
  jobId: string; // JobId
  exitCode: number;
  outputs: string[]; // CIDs
  stderr: string;
  stdout: string;
};

export default JobResult;