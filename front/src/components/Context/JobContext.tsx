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

type File = {
    cid: string;
    type: string;
    size: number;
}

type Job = {
    name: string;
    category: Category | null;
    image: string;
    generateInput: (file: File) => File;
}

type JobRequest = {
    usrInput?: File;
    input?: File;
    job?: Job;
    custom: boolean;
}

type JobContextType = {
    step: number;
    setStep: (step: number) => void;
    category: Category | null;
    setCategory: (category: Category | null) => void;
    job: Job | null;
    setJob: (job: Job | null) => void;
    jobRequest: JobRequest;
    setJobRequest: (jobRequest: JobRequest) => void;
    allCategories: Category[];
    allJobs: Job[];
    allFiles: File[];
}

/*
 * Step 0: Select Job
 * Step 1: Select Cid Input
 * Step 2: Submit Job
 * Step 3: Job Submitted
 */
const JobContext = createContext<JobContextType>({
    step: 0,
    category: null,
    job: null,
    jobRequest: {
        input: undefined,
        job: undefined,
        custom: false
    },
    setStep: () => {},
    setCategory: () => {},
    setJob: () => {},
    setJobRequest: () => {},
    allCategories: [],
    allJobs: [],
    allFiles: []
});

const useJobContext = () => React.useContext(JobContext);

const JobProvider = ({ children }: { children: React.ReactNode }) => {
    const [step, setStep] = useState(0);
    const [category, setCategory] = useState<Category | null>(null);
    const [job, setJob] = useState<Job | null>(null);
    const [jobRequest, setJobRequest] = useState<JobRequest>({
        input: undefined,
        job: undefined,
        custom: false
    });
    const [allCategories, setAllCategories] = useState<Category[]>([
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
    
    const mockGenerateInput = (file: File) => {
        return {
            cid: file.cid,
            type: 'img',
            size: file.size
        }
    }
    const [allJobs, setAllJobs] = useState<Job[]>([
        {
            name: 'Image to Image',
            category: allCategories[0],
            image: 'ubuntu:latest',
            generateInput: mockGenerateInput
        },
        {
            name: 'EasyOCR (OCR)',
            category: allCategories[0],
            image: 'ubuntu:latest',
            generateInput: mockGenerateInput
        },
        {
            name: 'Stable diffusion (GAN)',
            category: allCategories[2],
            image: 'ubuntu:latest',
            generateInput: mockGenerateInput
        },
        {
            name: 'Custom (Docker)',
            category: allCategories[4],
            image: 'ubuntu:latest',
            generateInput: mockGenerateInput
        },
        {
            name: 'Custom (Python)',
            category: allCategories[4],
            image: 'ubuntu:latest',
            generateInput: mockGenerateInput
        },
    ]);
    const [allFiles, setAllFiles] = useState<File[]>([
        {
            cid: "QmWXShtJXt6Mw3FH7hVCQvR56xPcaEtSj4YFSGjp2QxA4v",
            type: "img",
            size: 400000
          },
          {
            cid: "QmU7gJi6Bz3jrvbuVfB7zzXStLJrTHf6vWh8ZqkCsTGoRC",
            type: "json",
            size:  154555
          },
          {
            cid: "QmWXShtJXt6Mw3FH7hVCQvR56xPcaEtSj4YFSGjp2QxB4v",
            type: "txt",
            size: 234563
          },
    ]);

    return (
        <JobContext.Provider value={{
            step, setStep,
            category, setCategory,
            job, setJob,
            jobRequest, setJobRequest,
            allCategories, allJobs, allFiles
        }}>
            {children}
        </JobContext.Provider>
    )
}

export { useJobContext, JobProvider }
export type { Category, Job, File, JobRequest }