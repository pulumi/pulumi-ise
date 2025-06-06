// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.IdentityManagement.Outputs
{

    [OutputType]
    public sealed class GetActiveDirectoryGroupsByDomainGroupResult
    {
        /// <summary>
        /// Group name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Group SID
        /// </summary>
        public readonly string Sid;
        /// <summary>
        /// Group type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetActiveDirectoryGroupsByDomainGroupResult(
            string name,

            string sid,

            string type)
        {
            Name = name;
            Sid = sid;
            Type = type;
        }
    }
}
