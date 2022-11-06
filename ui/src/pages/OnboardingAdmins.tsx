import superclusterLogo from "../assets/superclusterLogo.svg";

function OnboardingAdmins() {
  return (
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        <h1 className="text-4xl font-bold text-white mb-10">
          ⭐️ Awesome! Who are your cluster’s admins?
        </h1>
        <p className="text-2xl text-l-slategray-50">
          Admins can adjust permissions, remove team members, and change your
          cluster’s settings. Make sure you trust them!
        </p>
      </div>
    </div>
  );
}

export default OnboardingAdmins;
