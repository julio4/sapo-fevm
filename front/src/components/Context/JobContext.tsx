import React, { createContext, useState } from "react";
import { IconType } from "react-icons";
import {
  FaImages,
  FaCogs,
  FaHdd,
  FaRegObjectGroup,
  FaPen,
} from "react-icons/fa";

type Category = {
  name: string;
  id: number;
  icon: IconType;
  color:
    | "red"
    | "green"
    | "blue"
    | "yellow"
    | "purple"
    | "pink"
    | "teal"
    | "cyan"
    | null;
};

type File = {
  cid: string;
  type: string;
  size: number | null;
};

type Job = {
  name: string;
  category: Category | null;
  image: string;
  inputTypes: string[];
  runParams: string[];
  desc: string;
  needInput: boolean;
  complexity: number;
};

type JobRequest = {
  usrInput?: File;
  input?: File;
  job?: Job;
  custom: boolean;
};

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
  setAllFiles: (files: File[]) => void;
  seeIncompatibleFiles: boolean;
  setSeeIncompatibleFiles: (see: boolean) => void;
};

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
    custom: false,
  },
  setStep: () => {},
  setCategory: () => {},
  setJob: () => {},
  setJobRequest: () => {},
  allCategories: [],
  allJobs: [],
  allFiles: [],
  setAllFiles: () => {},
  seeIncompatibleFiles: false,
  setSeeIncompatibleFiles: () => {},
});

const useJobContext = () => React.useContext(JobContext);

const JobProvider = ({ children }: { children: React.ReactNode }) => {
  const [step, setStep] = useState(0);
  const [category, setCategory] = useState<Category | null>(null);
  const [job, setJob] = useState<Job | null>(null);
  const [jobRequest, setJobRequest] = useState<JobRequest>({
    input: undefined,
    job: undefined,
    custom: false,
  });
  const [allCategories, setAllCategories] = useState<Category[]>([
    {
      id: 1,
      name: "Images processing",
      icon: FaImages,
      color: "blue",
    },
    {
      id: 2,
      name: "Data generation",
      icon: FaCogs,
      color: "red",
    },
    {
      id: 3,
      name: "Artifical Intelligence",
      icon: FaHdd,
      color: "teal",
    },
    {
      id: 4,
      name: "Script",
      icon: FaRegObjectGroup,
      color: "yellow",
    },
    {
      id: 5,
      name: "Custom",
      icon: FaPen,
      color: "purple",
    },
  ]);

  const [allJobs, setAllJobs] = useState<Job[]>([
    {
      name: "Pandas script",
      category: allCategories[4],
      image: "amancevice/pandas",
      inputTypes: ["text"],
      runParams: ["python3", "/inputs"],
      desc: "Execute python script with pandas (input: .py)",
      needInput: true,
      complexity: 2
    },
    {
      name: "Python Hello World",
      category: allCategories[3],
      image: "python:3.10-slim",
      inputTypes: [],
      runParams: ["python3", "-c", "print(\"Hello, world\")"],
      desc: "Print Hello world in python",
      needInput: false,
      complexity: 1
    },
    {
      name: "R script",
      category: allCategories[4],
      image: "r-base",
      inputTypes: ["text"],
      runParams: ["Rscript", "/inputs"],
      desc: "Execute R script (input: .r)",
      needInput: true,
      complexity: 2
    },
    {
      name: "Sparkov credit card",
      category: allCategories[1],
      image: "jsacex/sparkov-data-generation",
      inputTypes: [],
      runParams: [
        "python3", 
        "datagen.py", 
        "-n", 
        "1000", 
        "-o", 
        "../outputs", 
        "01-01-2023", 
        "10-01-2023"
      ],
      desc: "synthetic credit card transaction data generation using Sparkov",
      needInput: false,
      complexity: 2
    },
    {
      name: "Csv data cleaning",
      category: allCategories[4],
      image: "amancevice/pandas",
      inputTypes: ["text/csv"],
      runParams: ["python3", "-c", "import pandas as pd;df = pd.read_csv('./input');df.dropna(inplace=True);df.drop_duplicates(inplace=True);df.to_csv('./output/output.csv', index=False)"],
      desc: "Check for missing values, duplicates, and columns data types (input: .csv)",
      needInput: true,
      complexity: 1
    },
    {
      name: "Batch image resizing",
      category: allCategories[0],
      image: "dpokidov/imagemagick:7.1.0-47-ubuntu",
      inputTypes: ["text/html"],
      runParams: ["magick", "mogrify", "-resize", "100x100", "-quality", "100", "-path", "/outputs", "/inputs/*.jpg"],
      desc: "Resize a batch of jpg to 100x100 (input: ./*.jpg)",
      needInput: true,
      complexity: 3
    },
    {
      name: "Python script",
      category: allCategories[4],
      image: "python:3.10-slim",
      inputTypes: ["text"],
      runParams: ["python3", "/inputs"],
      desc: "Execute python script (input: .py)",
      needInput: true,
      complexity: 2
    },
    {
      name: "Solidity compiler",
      category: allCategories[4],
      image: "ethereum/solc:stable",
      inputTypes: ["text"],
      runParams: ["solc","-o", "/outputs", "--abi", "--bin", "/inputs"],
      desc: "Generate evm bytecode and ABI from solidity code (input: .sol)",
      needInput: true,
      complexity: 2
    },
    {
      name: "Stable Diffusion CPU",
      category: allCategories[2],
      image: "ghcr.io/bacalhau-project/examples/stable-diffusion-cpu:0.0.1",
      inputTypes: ["text/plain"],
      runParams: ["/bin/bash", "-c", "python demo.py --prompt \"$(cat /inputs)\" --output /outputs"],
      desc: "Stable diffusion image generation from prompt (input: .txt)",
      needInput: true,
      complexity: 5
    },
    {
      name: "Stable Diffusion GPU",
      category: allCategories[2],
      image: "ghcr.io/bacalhau-project/examples/stable-diffusion-gpu:0.0.1",
      inputTypes: ["text/plain"],
      runParams: ["/bin/bash", "-c", "python /main.py --o /outputs --p \"$(cat /inputs)\""],
      desc: "Stable diffusion image generation from prompt (input: .txt)",
      needInput: true,
      complexity: 5
    },
    {
      name: "Image to Text (EasyOCR)",
      category: allCategories[0],
      image: "jsacex/easyocr",
      inputTypes: ["img"],
      runParams: ["easyocr", "-l", "ch_sim", "en", "-f", "./inputs", "--detail=1", "--gpu=True"],
      desc: "(NOT WORKING: waiting for GPU support) Optical Character Recognition from jpg using EasyOCR (input: .jpg)",
      needInput: true,
      complexity: 5
    },
    {
      name: "Speech Recognition",
      category: allCategories[2],
      image: "jsacex/whisper",
      inputTypes: ["text"],
      runParams: ["python", "openai-whisper.py", "-p", "./inputs", "-o", "outputs"],
      desc: "Speech recognition using Whisper (input: .mp4|.mp3)",
      needInput: true,
      complexity: 5
    },
    {
      name: "Object Detection (YOLOv5)",
      category: allCategories[0],
      image: "ultralytics/yolov5:v6.2",
      inputTypes: ["img"],
      runParams: ["/bin/bash", "-c", "'find /inputs -type f -exec cp {} /outputs/yolov5s.pt \; ; python detect.py --weights /outputs/yolov5s.pt --source /datasets --project /outputs'"],
      desc: "(NOT WORKING: waiting for GPU support) Object detection with YOLOv5 (input: .jpg)",
      needInput: true,
      complexity: 5
    },
  ]);
  const [allFiles, setAllFiles] = useState<File[]>([]);
  const [seeIncompatibleFiles, setSeeIncompatibleFiles] = useState(false);

  return (
    <JobContext.Provider
      value={{
        step,
        setStep,
        category,
        setCategory,
        job,
        setJob,
        jobRequest,
        setJobRequest,
        allCategories,
        allJobs,
        allFiles,
        setAllFiles,
        seeIncompatibleFiles,
        setSeeIncompatibleFiles,
      }}
    >
      {children}
    </JobContext.Provider>
  );
};

export { useJobContext, JobProvider };
export type { Category, Job, File, JobRequest };
