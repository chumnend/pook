import { useState } from 'react';

import Step1BookName from './components/Step1BookName';
import Step2ManagePages from './components/Step2ManagePages';
import Step3EditPage from './components/Step3EditPages';
import Step4ValidateAndPublish from './components/Step4ValidateAndPublish';

import Header from '../../shared/Header';
import './BookCreationPage.module.css';

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
