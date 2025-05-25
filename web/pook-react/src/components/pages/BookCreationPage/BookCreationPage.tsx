import { useState } from 'react';

import Header from '../../shared/Header';
import './BookCreationPage.module.css';

export type Step1BookNameProps = {
  title: string;
  changeTitle: (event: React.ChangeEvent<HTMLInputElement>) => void;
  nextStep: () => void;
}

const Step1BookName = ({ title, changeTitle, nextStep }: Step1BookNameProps) => {
  return (
    <div>
      <h1>Let's think of a book title.</h1>
      <input 
        id="title-input"
        type="text" 
        placeholder='Book Title'
        value={title}
        onChange={changeTitle}
      />
      <button onClick={nextStep}>Next</button>
     </div>
  );
}

export type Step2ManagePagesProps = {
  title: string;
  previousStep: () => void;
  nextStep: () => void;
};

const Step2ManagePages = ({ title, previousStep, nextStep }: Step2ManagePagesProps) => {
  return (
    <div>
      <h1>Time to create some pages.</h1>
      <p>{title}</p>
      <button onClick={previousStep}>Prveious</button>
      <button onClick={nextStep}>Next</button>
     </div>
  );
}

export type Step3EditPageProps = {
  title: string;
  previousStep: () => void;
  nextStep: () => void;
};

const Step3EditPage = ({ title, previousStep, nextStep }: Step3EditPageProps) => {
  return (
    <div>
      <h1>Let's setup some content.</h1>
      <p>{title}</p>
      <button onClick={previousStep}>Prveious</button>
      <button onClick={nextStep}>Next</button>
     </div>
  );
}

export type Step4ValidateAndPublishProps = {
  title: string;
  previousStep: () => void;
  handlePublish: () => void;
};

const Step4ValidateAndPublish = ({ title, previousStep, handlePublish }: Step4ValidateAndPublishProps) => {
  return (
    <div>
      <h1>Validate your current story.</h1>
      <p>{title}</p>
      <button onClick={previousStep}>Prveious</button>
      <button onClick={handlePublish}>Publish</button>
     </div>
  );
}

const BookCreationPage = () => {
  const [step, setStep] = useState<number>(1);
  const [title, setTitle] = useState<string>('');

  const previousStep = () => {
    setStep(currentStep => currentStep - 1);
  }

  const nextStep = () => {
    setStep(currentStep => currentStep + 1);
  }

  const changeTitle = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };
  
  const handlePublish = () => {
    return null;
  }

  return (
    <div>
     <Header />
      {step === 1 && (
        <Step1BookName 
          title={title}
          changeTitle={changeTitle}
          nextStep={nextStep}
        />
      )}
      {step === 2 && (
        <Step2ManagePages 
          title={title}
          previousStep={previousStep}
          nextStep={nextStep}
        />
      )}
      {step === 3 && (
        <Step3EditPage 
          title={title}
          previousStep={previousStep}
          nextStep={nextStep}
        />
      )}
      {step === 4 && (
        <Step4ValidateAndPublish
          title={title}
          previousStep={previousStep}
          handlePublish={handlePublish}
        />
      )}
    </div>
  );
}

export default BookCreationPage;
