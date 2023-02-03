import React, { createContext, useState } from "react";
import { IconType } from 'react-icons';

type Category = {
    name: string;
    id: number;
    icon: IconType;
}

type Job = {
    name: string;
    category: Category | null;
}

type JobRequest = {
    image: string;
    input: string;
    cid: string;
}

type JobContextType = {
    step: number;
    setStep: (step: number) => void;
    category: Category | null;
    setCategory: (category: Category) => void;
    job: Job | null;
    setJob: (job: Job) => void;
    jobRequest: JobRequest | null;
    setJobRequest: (jobRequest: JobRequest) => void;
}

/*
 * Step 1: Select Job
 * Step 2: Select Cid Input
 * Step 3: Submit Job
 * Step 4: Job Submitted
 */
const JobContext = createContext<JobContextType>({
    step: 0,
    category: null,
    job: null,
    jobRequest: null,
    setStep: () => {},
    setCategory: () => {},
    setJob: () => {},
    setJobRequest: () => {},
});

const useJobContext = () => React.useContext(JobContext);

const JobProvider = ({ children }: { children: React.ReactNode }) => {
    const [step, setStep] = useState(0);
    const [category, setCategory] = useState<Category | null>(null);
    const [job, setJob] = useState<Job | null>(null);
    const [jobRequest, setJobRequest] = useState<JobRequest | null>(null);
    return (
        <JobContext.Provider value={{
            step, setStep,
            category, setCategory,
            job, setJob,
            jobRequest, setJobRequest
        }}>
            {children}
        </JobContext.Provider>
    )
}

export { useJobContext, JobProvider }
export type { Category }