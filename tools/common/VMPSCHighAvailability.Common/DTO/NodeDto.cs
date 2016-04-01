﻿/*
using System;

namespace VMPSCHighAvailability.Common.DTO
{
    public class NodeDto
    {
        /// </summary>
        public string Domain { get; set; }
        /// <summary>
        /// Site name
        /// </summary>
        public string Sitename { get; set; }
        /// <summary>
        /// Name of the node
        /// </summary>
        public string Name { get; set; }
        /// <summary>
        /// IP Address
        /// </summary>
        public string Ip { get; set; }
        /// <summary>
        /// Gets or sets the last heart beat.
        /// </summary>
        /// <value>The last heart beat.</value>
        public DateTime? LastHeartBeat { get; set; }
        /// <summary>
        /// Gets or sets the type of the node.
        /// </summary>
        /// <value>The type of the node.</value>
        public NodeType NodeType { get; set; }
        /// <summary>
        /// Is the node active or inactive
        /// </summary>
        public bool Active { get; set; }
    }
}