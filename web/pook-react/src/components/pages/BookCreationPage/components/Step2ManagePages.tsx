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

export default Step2ManagePages;
