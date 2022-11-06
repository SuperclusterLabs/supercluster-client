import superclusterLogo from "../assets/superclusterLogo.svg";
import ButtonPrimary from "../components/ButtonPrimary";
import ButtonSecondary from "../components/ButtonSecondary";

function OnboardingInvite() {
  function generateShareLink() {
    console.log("Generating share link");
  }

  function skip() {
    console.log("skipping");
  }
  return (
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        <h1 className="text-4xl font-bold text-white mb-10">
          Perfect! 🤝🏼 Lastly, invite your team members!
        </h1>
        <p className="text-2xl text-l-slategray-50">
          Share a link to your cluster with your team members. They’ll be able
          to easily join your cluster.
        </p>
        <ButtonPrimary onClick={generateShareLink} text="Generate share link" />
        <ButtonSecondary onClick={skip} text="Skip for now" />
      </div>
    </div>
  );
}

export default OnboardingInvite;
