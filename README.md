<h2>Voucher service</h2>

<h3>Project Structure</h3>
<ul>
    <li>api</li>
        <div>
            This package contains proto files and handler API documentation. 
        </div>
    <li>cmd</li>
        <div>
            Service Entry point and application configs available here. 
        </div>
    <li>configs</li>
        <div>
            This package contains default configs of application
        </div>
    <li>deploy</li>
        <div>
            Deployment scripts implemented here.
        </div>
    <li>internal</li>
        <div>
            <div>
            Core files of service implemented here. Internal section contains two packages. app and pkg.
            <br>
            pkg contains shared codes between all apps like database interface and entity.
            <br>
            app contains implementation of each rpc services, HTTP handlers and service workers.
        </div>
        </div>
    <li>pkg</li>
        <div>
            Shared packages implemented there.
        </div>
        <p>
            Note: It's better to put all shared packages in a seperated repository
        </p>
    <li>script</li>
        <div>
            bash script of service contains here. these scripts are deployment script, running of program, proto generator, ci/cd scripts etc...
        </div>

</ul>

<h3>Improvement notes</h3>

<ul>
    <li>
        <div>
            Using redis lock in applying voucher for every wallet.
        </div>
    </li>
    <li>
        <div>
            Separating HttpAPIs into a gateway service.
        </div>
    </li>
    <li>
        <div>
            Move all shared packages in into a shared project.
        </div>
    </li>
</ul>