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
      <button onClick={previousStep}>Previous</button>
      <button onClick={nextStep}>Next</button>
     </div>
  );
}

export default Step3EditPage;
