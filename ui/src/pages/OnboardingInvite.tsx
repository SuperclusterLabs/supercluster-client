import ButtonPrimary from "../components/ButtonPrimary";
import ButtonSecondary from "../components/ButtonSecondary";
import { useNavigate } from "react-router-dom";

function OnboardingInvite() {
  const navigate = useNavigate();

  function generateShareLink() {
    console.log("Generating share link");
  }

  function skip() {
    navigate("/cluster");
  }
  return (
    <div className="text-l-slateblue-700 mt-6">
      <h1 className="text-4xl font-bold">
        Perfect! 🤝 Lastly, invite your team members!
      </h1>
      <p className="text-xl mt-4">
        Share a link to your cluster with your team members. They’ll be able
        to easily join your cluster.
      </p>
      <ButtonPrimary onClick={generateShareLink} text="Generate share link" />
      <ButtonSecondary onClick={skip} text="Skip for now" />
    </div>
  );
}

export default OnboardingInvite;
