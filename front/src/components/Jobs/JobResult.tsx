type JobResult = {
  id: number;
  address: `0x${string}`;
  exitCode: number;
  outputs: string[]; // CIDs
  stderr: string;
  stdout: string;
};

export default JobResult;