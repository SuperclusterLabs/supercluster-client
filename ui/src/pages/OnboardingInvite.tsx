import ButtonPrimary from "../components/ButtonPrimary";
import ButtonSecondary from "../components/ButtonSecondary";
import { useNavigate } from "react-router-dom";

function OnboardingInvite() {
  const navigate = useNavigate();

  function generateShareLink() {
    console.log("Generating share link");
  }

  function skip() {
    navigate("Files");
  }
  return (
    <div className="text-l-slateblue-700">
      <h1 className="text-2xl font-bold">
        Perfect! 🤝 Lastly, invite your team members!
      </h1>
      <p className="text-lg">
        Share a link to your cluster with your team members. They’ll be able
        to easily join your cluster.
      </p>
      <ButtonPrimary onClick={generateShareLink} text="Generate share link" />
      <ButtonSecondary onClick={skip} text="Skip for now" />
    </div>
  );
}

export default OnboardingInvite;
