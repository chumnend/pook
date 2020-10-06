import styled from 'styled-components';

export const PageContent = styled.div`
  width: 100%;
  height: 100%;
  min-height: 768px;
  padding: 2rem;
`;

export const StyledForm = styled.form`
  width: 50%;
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  border: 1px solid #000;
  & p {
    text-align: center;
  }
`;

export const StyledFormHeader = styled.div`
  margin-bottom: 1rem;
  text-align: center;
  & h2 {
    font-size: 1.5rem;
  }
  & p {
    margin: 1rem 0;
    background: red;
    color: #000;
  }
`;

export const StyledFormGroup = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
  & label {
    margin-bottom: 0.3rem;
    font-size: 1rem;
  }
  & input {
    padding: 0.8rem;
  }
  & small {
    font-size: 0.8rem;
    color: red;
  }
`;

export const StyledButton = styled.button`
  width: 100%;
  margin: 1rem auto;
  padding: 0.8rem 1rem;
  font-size: 1rem;
  cursor: pointer;
`;
