import React, { createContext, useState } from "react";
import { IconType } from 'react-icons';
import {
    FaImages,
    FaCogs,
    FaHdd,
    FaRegObjectGroup,
    FaPen,
} from 'react-icons/fa';

type Category = {
    name: string;
    id: number;
    icon: IconType;
    color: "red" | "green" | "blue" | "yellow" | "purple" | "pink" | "teal" | "cyan" | null;
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
    setCategory: (category: Category | null) => void;
    job: Job | null;
    setJob: (job: Job | null) => void;
    jobRequest: JobRequest | null;
    setJobRequest: (jobRequest: JobRequest) => void;
    allCategories: Category[];
    allJobs: Job[];
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
    allCategories: [],
    allJobs: []
});

const useJobContext = () => React.useContext(JobContext);

const JobProvider = ({ children }: { children: React.ReactNode }) => {
    const [step, setStep] = useState(0);
    const [category, setCategory] = useState<Category | null>(null);
    const [job, setJob] = useState<Job | null>(null);
    const [jobRequest, setJobRequest] = useState<JobRequest | null>(null);
    const [allCategories, setAllCategories] = useState([
        {
            id: 1,
            name: 'Images processing',
            icon: FaImages,
            color: "blue"
        },
        {
            id: 2,
            name: 'Data generation',
            icon: FaCogs,
            color: "red"
        },
        {
            id: 3,
            name: 'Artifical Intelligence',
            icon: FaHdd,
            color: "teal"
        },
        {
            id: 4,
            name: 'Simulation',
            icon: FaRegObjectGroup,
            color: "yellow"
        },
        {
            id: 5,
            name: 'Custom',
            icon: FaPen,
            color: "purple"
        }
    ]);
    const [allJobs, setAllJobs] = useState([
        {
            name: 'Image to Image',
            category: allCategories[0]
        },
        {
            name: 'EasyOCR (OCR)',
            category: allCategories[0]
        },
        {
            name: 'Stable diffusion (GAN)',
            category: allCategories[2]
        },
        {
            name: 'Custom (Docker)',
            category: allCategories[4]
        },
        {
            name: 'Custom (Python)',
            category: allCategories[4]
        },
    ]);
    return (
        <JobContext.Provider value={{
            step, setStep,
            category, setCategory,
            job, setJob,
            jobRequest, setJobRequest,
            allCategories, allJobs
        }}>
            {children}
        </JobContext.Provider>
    )
}

export { useJobContext, JobProvider }
export type { Category, Job }